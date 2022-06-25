package config

import (
	"github.com/spf13/viper"
)

var (
	cfg *config
)

type config struct {
	Server     ServerConfig
	DataSource DataSourceConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DataSourceConfig struct {
	URL string
}

func init() {
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "8080")
}

func LoadConfig() error {
	viper.SetConfigName("properties")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./resources/")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.Server = ServerConfig{
		Host: viper.GetString("server.host"),
		Port: viper.GetString("server.port"),
	}

	cfg.DataSource = DataSourceConfig{
		URL: viper.GetString("datasource.url"),
	}

	return nil
}

func GetServer() ServerConfig {
	return cfg.Server
}

func GetDataSource() DataSourceConfig {
	return cfg.DataSource
}
