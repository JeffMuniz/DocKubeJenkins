// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package oracle

import (
	"os"

	"github.com/godror/godror"
	"github.com/godror/godror/dsn"
)

// GetOracleConnectionDetails return a valid SID to use for testing
func GetOracleConnectionDetails(host string) string {
	params := godror.ConnectionParams{
		CommonParams: dsn.CommonParams{
			Username: GetOracleEnvUsername(),
			Password: dsn.NewPassword((GetOracleEnvPassword())),
		},
		ConnParams: dsn.ConnParams{
			IsSysDBA: true,
		},
	}

	return params.StringWithPassword()
}

// GetOracleEnvServiceName returns the service name to use with Oracle testing server or the value of the environment variable ORACLE_SERVICE_NAME if not empty
func GetOracleEnvServiceName() string {
	serviceName := os.Getenv("ORACLE_SERVICE_NAME")

	if len(serviceName) == 0 {
		serviceName = "ORCLCDB.localdomain"
	}
	return serviceName
}

// GetOracleEnvUsername returns the username to use with Oracle testing server or the value of the environment variable ORACLE_USERNAME if not empty
func GetOracleEnvUsername() string {
	username := os.Getenv("ORACLE_USERNAME")

	if len(username) == 0 {
		username = "sys"
	}
	return username
}

// GetOracleEnvUsername returns the port of the Oracle server or the value of the environment variable ORACLE_PASSWORD if not empty
func GetOracleEnvPassword() string {
	password := os.Getenv("ORACLE_PASSWORD")

	if len(password) == 0 {
		password = "Oradoc_db1" // #nosec
	}
	return password
}
