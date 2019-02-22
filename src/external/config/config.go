package config

import "github.com/spf13/viper"

func Load() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}

	return c
}

type Config struct {
	DB struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
	}
}
