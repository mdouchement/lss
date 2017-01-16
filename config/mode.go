package config

import (
	"fmt"
	"os"
)

const (
	// EnvironmentVariableName is the name of the environment variable that defines the application mode.
	EnvironmentVariableName = "ECHO_ENV"
	// Development environment name
	Development = "development"
	// Test environment name
	Test = "test"
	// Production environment name
	Production = "production"
)

var envMode = Development

func init() {
	if e := os.Getenv(EnvironmentVariableName); e != "" {
		switch e {
		case Development, Test, Production:
			envMode = e
		default:
			panic(fmt.Sprintf("Bad %s value: %s", EnvironmentVariableName, e))
		}
	}
}

// Env returns the current application environment.
func Env() string {
	return envMode
}

// SetEnv sets the application environment.
func SetEnv(environment string) {
	envMode = environment
}
