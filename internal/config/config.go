package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" envDefault:"development" envRequired:"true"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env:"HTTP_ADDRESS" envDefault:":8080"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" envDefault:"5s"`
	IdleTimeout time.Duration `yaml:"idleTimeout" env:"HTTP_IDLE_TIMEOUT" envDefault:"60s"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatal(err)
	}

	return &config
}
