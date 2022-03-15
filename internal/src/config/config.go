package config

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/joomcode/errorx"
	"github.com/spf13/viper"
)

type Configs struct {
	AppName       string
	ServerHost    string
	ServerMonitor string
	DatabaseDSN   string
	DatabaseName  string
}

var config = viper.New()

var onceConfigs sync.Once
var configs Configs

func Init() *viper.Viper {
	config.AddConfigPath("internal/src/config/")
	config.SetConfigName("configuration")
	config.SetConfigType("yml")

	setConfigDefaults()

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error was produced
			errorx := errorx.Decorate(err, "Error reading config file")
			log.Println(errorx)
			// nrlog.Logger(context.Background()).Fatal(errorx)
		}
	}

	return config
}

func setConfigDefaults() {
	config.SetDefault("app.name", "hexagonal-template")
	config.SetDefault("server.host", "0.0.0.0:8080")

	//Postgres config
	config.SetDefault("database.host", "localhost")
	config.SetDefault("database.database", "dummy_db")
	config.SetDefault("database.port", 5432)
	config.SetDefault("database.user", "dummy_db")
	config.SetDefault("database.password", "dummy_db")

	//Log
	config.SetDefault("logger.level", "info")

	config.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	config.AutomaticEnv()
}

func GetConfig() Configs {
	onceConfigs.Do(func() {
		viperConfig := Init()

		databaseDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s connect_timeout=5",
			viperConfig.Get("database.host"),
			viperConfig.Get("database.user"),
			viperConfig.Get("database.password"),
			viperConfig.Get("database.database"),
			viperConfig.GetString("database.port"))

		configs = Configs{
			DatabaseDSN:   databaseDSN,
			DatabaseName:  viperConfig.GetString("database.database"),
			AppName:       viperConfig.GetString("app.name"),
			ServerHost:    viperConfig.GetString("server.host"),
			ServerMonitor: viperConfig.GetString("server.monitor"),
		}
	})

	return configs
}
