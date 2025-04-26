package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	App string `env:"APP" env-default:"dev" env-required:"true"`
	Storage
	HTTPServer
}

type HTTPServer struct {
	Address     string        `env:"HTTP_ADDRESS" env-default:":8080"`
	Timeout     time.Duration `env:"HTTP_TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `env:"HTTP_IDLE_TIMEOUT" env-default:"60s"`
}

type Storage struct {
	Host   string `env:"STORAGE_HOST" env-default:"localhost" env-required:"true"`
	Port   string `env:"STORAGE_PORT" env-default:"8080" env-required:"true"`
	User   string `env:"STORAGE_USER" env-default:"development" env-required:"true"`
	Pass   string `env:"STORAGE_PASS" env-default:"development" env-required:"true"`
	DBName string `env:"STORAGE_DB_NAME" env-default:"urlShortener" env-required:"true"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	var config Config

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("config: %#v\n", config)

	return &config
}
