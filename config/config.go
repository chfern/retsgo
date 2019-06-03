package config

import (
	"github.com/fernandochristyanto/retsgo/config/keys"
	"os"
	"sync"
)

// Conf holds the initialized config struct
var (
	conf        Config
	once        sync.Once
	initialized bool
)

// GetConf initializes the config fields according to app env
func GetConf() Config {
	appEnv := os.Getenv(keys.Env.AppEnv)
	once.Do(func() {
		if initialized == false {
			conf = getConfigLoader(appEnv).Load()
			conf.General = loadGeneralConf()
			initialized = true
		}
	})

	return conf
}
