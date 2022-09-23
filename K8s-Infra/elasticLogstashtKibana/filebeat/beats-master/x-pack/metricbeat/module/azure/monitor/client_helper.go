// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package monitor

import (
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-06-01/insights"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-03-01/resources"

	"github.com/pkg/errors"

	"github.com/elastic/beats/x-pack/metricbeat/module/azure"
)

// mapMetric should validate and map the metric related configuration to relevant azure monitor api parameters
func mapMetric(client *azure.Client, metric azure.MetricConfig, resource resources.GenericResource) ([]azure.Metric, error) {
	var metrics []azure.Metric
	// get all metrics supported by the namespace provided
	metricDefinitions, err := client.AzureMonitorService.GetMetricDefinitions(*resource.ID, metric.Namespace)
	if err != nil {
		return nil, errors.Wrapf(err, "no metric definitions were found for resource %s and namespace %s.", *resource.ID, metric.Namespace)
	}
	if len(*metricDefinitions.Value) == 0 {
		return nil, errors.Errorf("no metric definitions were found for resource %s and namespace %s.", *resource.ID, metric.Namespace)
	}

	// validate metric names
	// check if all metric names are selected (*)
	var supportedMetricNames []string
	var unsupportedMetricNames []string
	if strings.Contains(strings.Join(metric.Name, " "), "*") {
		for _, definition := range *metricDefinitions.Value {
			supportedMetricNames = append(supportedMetricNames, *definition.Name.Value)
		}
	} else {
		// verify if configured metric names are valid, return log error event for the invalid ones, map only  the valid metric names
		supportedMetricNames, unsupportedMetricNames = filterSConfiguredMetrics(metric.Name, *metricDefinitions.Value)
		if len(unsupportedMetricNames) > 0 {
			return nil, errors.Errorf("the metric names configured  %s are not supported for the resource %s and namespace %s",
				strings.Join(unsupportedMetricNames, ","), *resource.ID, metric.Namespace)
		}
	}
	if len(supportedMetricNames) == 0 {
		return nil, errors.Errorf("the metric names configured : %s are not supported for the resource %s and namespace %s ", strings.Join(metric.Name, ","), *resource.ID, metric.Namespace)
	}
	//validate aggregations and filter on supported ones
	var supportedAggregations []string
	var unsupportedAggregations []string
	metricGroups := make(map[string][]insights.MetricDefinition)
	metricDefs := getMetricDefinitionsByNames(*metricDefinitions.Value, supportedMetricNames)

	if len(metric.Aggregations) == 0 {
		for _, metricDef := range metricDefs {
			metricGroups[string(metricDef.PrimaryAggregationType)] = append(metricGroups[string(metricDef.PrimaryAggregationType)], metricDef)
		}
	} else {
		supportedAggregations, unsupportedAggregations = filterAggregations(metric.Aggregations, metricDefs)
		if len(unsupportedAggregations) > 0 {
			return nil, errors.Errorf("the aggregations configured : %s are not supported for some of the metrics selected %s ",
				strings.Join(unsupportedAggregations, ","), strings.Join(supportedMetricNames, ","))
		}
		if len(supportedAggregations) == 0 {
			return nil, errors.Errorf("no aggregations were found based on the aggregation values configured or supported between the metrics : %s",
				strings.Join(supportedMetricNames, ","))
		}
		key := strings.Join(supportedAggregations, ",")
		metricGroups[key] = append(metricGroups[key], metricDefs...)
	}
	// map dimensions
	var dim []azure.Dimension
	if len(metric.Dimensions) > 0 {
		for _, dimension := range metric.Dimensions {
			dim = append(dim, azure.Dimension{Name: dimension.Name, Value: dimension.Value})
		}
	}
	for key, metricGroup := range metricGroups {
		var metricNames []string
		for _, metricName := range metricGroup {
			metricNames = append(metricNames, *metricName.Name.Value)
		}
		metrics = append(metrics, client.CreateMetric(resource, metric.Namespace, metricNames, key, dim, metric.Timegrain))
	}
	return metrics, nil
}

// filterSConfiguredMetrics will filter out any unsupported metrics based on the namespace selected
func filterSConfiguredMetrics(selectedRange []string, allRange []insights.MetricDefinition) ([]string, []string) {
	var inRange []string
	var notInRange []string
	var allMetrics string
	for _, definition := range allRange {
		allMetrics += *definition.Name.Value + " "
	}
	for _, name := range selectedRange {
		if strings.Contains(allMetrics, name) {
			inRange = append(inRange, name)
		} else {
			notInRange = append(notInRange, name)
		}

	}
	return inRange, notInRange
}

// filterAggregations will filter out any unsupported aggregations based on the metrics selected
func filterAggregations(selectedRange []string, metrics []insights.MetricDefinition) ([]string, []string) {
	var difference []string
	var supported = []string{"Average", "Maximum", "Minimum", "Count", "Total"}

	for _, metric := range metrics {
		var metricSupported []string
		for _, agg := range *metric.SupportedAggregationTypes {
			metricSupported = append(metricSupported, string(agg))
		}
		supported, _ = intersections(metricSupported, supported)
	}
	if len(selectedRange) != 0 {
		supported, difference = intersections(supported, selectedRange)
	}
	return supported, difference
}

// filter is a helper method, will filter out strings not part of a slice
func filter(src []string) (res []string) {
	for _, s := range src {
		newStr := strings.Join(res, " ")
		if !strings.Contains(newStr, s) {
			res = append(res, s)
		}
	}
	return
}

// intersections is a helper method, will compare 2 slices and return their intersection and difference records
func intersections(supported, selected []string) ([]string, []string) {
	var intersection []string
	var difference []string
	str1 := strings.Join(filter(supported), " ")
	for _, s := range filter(selected) {
		if strings.Contains(str1, s) {
			intersection = append(intersection, s)
		} else {
			difference = append(difference, s)
		}
	}
	return intersection, difference
}

// getMetricDefinitionsByNames is a helper method, will compare 2 slices and return their intersection
func getMetricDefinitionsByNames(metricDefs []insights.MetricDefinition, names []string) []insights.MetricDefinition {
	var metrics []insights.MetricDefinition
	for _, def := range metricDefs {
		for _, supportedName := range names {
			if *def.Name.Value == supportedName {
				metrics = append(metrics, def)
			}
		}
	}
	return metrics
}
