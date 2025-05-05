package cleanenv

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	App string `env:"APP" env-bootstrap:"dev" env-required:"true"`
	Storage
	HttpServer
}

type HttpServer struct {
	Address     string        `env:"HTTP_ADDRESS" env-bootstrap:":8080"`
	Timeout     time.Duration `env:"HTTP_TIMEOUT" env-bootstrap:"5s"`
	IdleTimeout time.Duration `env:"HTTP_IDLE_TIMEOUT" env-bootstrap:"60s"`
}

type Storage struct {
	Host     string `env:"STORAGE_HOST" env-bootstrap:"localhost" env-required:"true"`
	Port     string `env:"STORAGE_PORT" env-bootstrap:"8080" env-required:"true"`
	User     string `env:"STORAGE_USER" env-bootstrap:"development" env-required:"true"`
	Password string `env:"STORAGE_PASS" env-bootstrap:"development" env-required:"true"`
	DbName   string `env:"STORAGE_DB_NAME" env-bootstrap:"urlShortener" env-required:"true"`
}

func Factory() *Config {
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
