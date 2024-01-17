package code

var ViperConfiguration = `package config

import "github.com/spf13/viper"

func NewViper(in string) *viper.Viper {
	v := viper.New()

	switch in {
	case "prod":
		in = "config.prod"
	case "dev":
		in = "config.dev"
	default:
		in = "config.dev"
	}

	v.SetConfigName(in)
	v.SetConfigType("yaml")
	v.AddConfigPath("../")
	v.AddConfigPath("./")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	return v
}`
