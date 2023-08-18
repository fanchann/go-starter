package code

import (
	"encoding/json"
	"fmt"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"

	"github.com/fanchann/go-starter/helpers"
)

type DatabaseConfig struct {
	Driver       string `toml:"driver" yaml:"driver" json:"driver"`
	Host         string `toml:"host" yaml:"host" json:"host"`
	Port         int    `toml:"port" yaml:"port" json:"port"`
	Username     string `toml:"username" yaml:"username" json:"username"`
	Password     string `toml:"password" yaml:"password" json:"password"`
	Name         string `toml:"name" yaml:"name" json:"name"`
	SSLMode      string `toml:"sslmode" yaml:"sslmode" json:"sslmode"`
	MaxIdleConn  int    `toml:"max_idle_conn" yaml:"max_idle_conn" json:"max_idle_conn"`
	MaxOpenConn  int    `toml:"max_open_conn" yaml:"max_open_conn" json:"max_open_conn"`
	LifeTimeConn int    `toml:"life_time_conn" yaml:"life_time_conn" json:"life_time_conn"`
}

type AppConfig struct {
	Database DatabaseConfig `toml:"database" yaml:"database" json:"database"`
}

func WriteAppConfiguration(extension, host, driver, username, password, dbname string, port int) string {
	config := AppConfig{
		Database: DatabaseConfig{
			Driver:       driver,
			Host:         host,
			Port:         port,
			Username:     username,
			Password:     password,
			Name:         dbname,
			SSLMode:      "disable",
			MaxIdleConn:  30,
			MaxOpenConn:  20,
			LifeTimeConn: 40,
		},
	}

	switch extension {
	case "toml":
		tomlCfg, err := toml.Marshal(config)
		helpers.ErrorWithLog(err)
		return string(tomlCfg)
	case "json":
		jsonCfg, err := json.MarshalIndent(config, "", "  ")
		helpers.ErrorWithLog(err)
		return string(jsonCfg)
	case "yaml":
		yamlCfg, err := yaml.Marshal(config)
		helpers.ErrorWithLog(err)
		return string(yamlCfg)
	case "env":
		environment := func(config DatabaseConfig) string {
			var content string

			entries := []struct {
				key   string
				value interface{}
			}{
				{"db_driver", config.Driver},
				{"db_host", config.Host},
				{"db_username", config.Username},
				{"db_password", config.Password},
				{"db_port", config.Port},
				{"db_name", config.Name},
				{"db_sslmode", config.SSLMode},
			}

			for _, entry := range entries {
				content += fmt.Sprintf("%s = \"%v\"\n", entry.key, entry.value)
			}

			return content
		}(config.Database)
		return environment
	default:
		tomlCfg, err := toml.Marshal(config)
		helpers.ErrorWithLog(err)
		return string(tomlCfg)
	}
}
