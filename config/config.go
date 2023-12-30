package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	LogLevel string
	Server
	Redis  Redis
	Logger Logger
}

type Server struct {
	ApiPort int
	Mode    string
}

type Logger struct {
	Encoding string
}

type Redis struct {
	Addr     string
	Username string
	Password string
	DB       int
}

func Load(path string) *Config {
	var c Config

	v := viper.New()

	v.AddConfigPath(path)
	v.SetConfigName("config.yml")
	v.SetConfigType("yml")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("config.ReadInConfig: %v", err)
	}

	if err := v.Unmarshal(&c); err != nil {
		log.Fatalf("config.Load.Unmarshal: %v", err)
	}

	return &c
}
