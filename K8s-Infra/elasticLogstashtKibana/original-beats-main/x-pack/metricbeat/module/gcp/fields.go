// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package gcp

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "gcp", asset.ModuleFieldsPri, AssetGcp); err != nil {
		panic(err)
	}
}

// AssetGcp returns asset data.
// This is the base64 encoded zlib format compressed contents of module/gcp.
func AssetGcp() string {
	return "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"
}
