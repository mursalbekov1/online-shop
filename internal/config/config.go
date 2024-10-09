package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	env        string     `yaml:env`
	HttpServer HttpServer `yaml:"http-server"`
}

type HttpServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func MustLoad() *Config {
	configPath := "config/config.yaml"

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Loaded config: %+v\n", cfg) // Вывод для отладки

	return &cfg
}
