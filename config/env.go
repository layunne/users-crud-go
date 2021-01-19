package config

import (
	"os"
)

type Env interface {
	WebServerPort() string
}

func NewEnv() Env {
	return &env{}
}

type env struct {
}

func (env) WebServerPort() string {
	return getString("SERVER_PORT", "8080")
}

func getString(name string, defaultValue string) string {
	envValue := os.Getenv(name)

	if envValue == "" {

		return defaultValue
	}

	return envValue
}
