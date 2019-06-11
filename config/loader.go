package config

import (
	"github.com/fernandochristyanto/retsgo/config/keys"
	"os"
	"strconv"
)

// General keeps all the config using the same env key
type General struct {
	AppPort              string
	ListenerReadTimeout  int
	ListenerWriteTimeout int
	JWTSecretKey         string
	JWTTimeoutMins       int
}

// Config keeps all variables (env) in an app
type Config struct {
	*General
	DatabaseHost     string
	DatabaseUsername string
	DatabasePassword string
	DatabasePort     string
	DatabaseName     string
	DatabaseSSL      string
	AppEnv           string
}

type baseConfig interface {
	Load() Config
}
type developmentConfig struct{}
type testConfig struct{}
type productionConfig struct{}

func loadGeneralConf() *General {
	listenerReadTimeout, _ := strconv.Atoi(os.Getenv(keys.Env.ListenerReadTimeout))
	listenerWriteTimeout, _ := strconv.Atoi(os.Getenv(keys.Env.ListenerWriteTimeout))
	jwtTimeoutMins, _ := strconv.Atoi(os.Getenv(keys.Env.JWTTimeoutMins))
	return &General{
		AppPort:              os.Getenv(keys.Env.AppPort),
		ListenerReadTimeout:  listenerReadTimeout,
		ListenerWriteTimeout: listenerWriteTimeout,
		JWTSecretKey:         os.Getenv(keys.Env.JWTSecretKey),
		JWTTimeoutMins:       jwtTimeoutMins,
	}
}

// Load loads config for development
func (developmentConfig) Load() Config {
	return Config{
		DatabaseHost:     os.Getenv(keys.Env.DatabaseHost),
		DatabaseUsername: os.Getenv(keys.Env.DatabaseUsername),
		DatabasePassword: os.Getenv(keys.Env.DatabasePassword),
		DatabasePort:     os.Getenv(keys.Env.DatabasePort),
		DatabaseName:     os.Getenv(keys.Env.DatabaseName),
		DatabaseSSL:      os.Getenv(keys.Env.DatabaseSSL),
		AppEnv:           "development",
	}
}

// Load loads config for testing
func (testConfig) Load() Config {
	return Config{
		DatabaseHost:     os.Getenv(keys.Env.TestDatabaseHost),
		DatabaseUsername: os.Getenv(keys.Env.TestDatabaseUsername),
		DatabasePassword: os.Getenv(keys.Env.TestDatabasePassword),
		DatabasePort:     os.Getenv(keys.Env.TestDatabasePort),
		DatabaseName:     os.Getenv(keys.Env.TestDatabaseName),
		DatabaseSSL:      os.Getenv(keys.Env.TestDatabaseSSL),
		AppEnv:           "test",
	}
}

// Load loads config for production
func (productionConfig) Load() Config {
	return Config{
		DatabaseHost:     os.Getenv(keys.Env.DatabaseHost),
		DatabaseUsername: os.Getenv(keys.Env.DatabaseUsername),
		DatabasePassword: os.Getenv(keys.Env.DatabasePassword),
		DatabasePort:     os.Getenv(keys.Env.DatabasePort),
		DatabaseName:     os.Getenv(keys.Env.DatabaseName),
		DatabaseSSL:      os.Getenv(keys.Env.DatabaseSSL),
		AppEnv:           "production",
	}
}

// GetConfigLoader returns the overloaded form of BaseConfig to load Config
func getConfigLoader(env string) baseConfig {
	if env == "development" {
		return new(developmentConfig)
	} else if env == "test" {
		return new(testConfig)
	} else if env == "production" {
		return new(productionConfig)
	}
	return nil
}
