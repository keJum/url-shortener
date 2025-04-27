package cleanenv

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
	configInterface "url-shortener/internal/config"
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
	Host   string `env:"STORAGE_HOST" env-bootstrap:"localhost" env-required:"true"`
	Port   string `env:"STORAGE_PORT" env-bootstrap:"8080" env-required:"true"`
	User   string `env:"STORAGE_USER" env-bootstrap:"development" env-required:"true"`
	Pass   string `env:"STORAGE_PASS" env-bootstrap:"development" env-required:"true"`
	DbName string `env:"STORAGE_DB_NAME" env-bootstrap:"urlShortener" env-required:"true"`
}

func Factory() configInterface.Config {
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

func (c *Config) GetApp() string {
	return c.App
}

func (c *Config) GetHTTPServer() configInterface.HTTPServer {
	return &c.HttpServer
}

func (h *HttpServer) GetAddress() string {
	return h.Address
}

func (h *HttpServer) GetTimeout() time.Duration {
	return h.Timeout
}

func (h *HttpServer) GetIdleTimeout() time.Duration {
	return h.IdleTimeout
}

func (c *Config) GetStorage() configInterface.Storage {
	return &c.Storage
}

func (s *Storage) GetHost() string {
	return s.Host
}

func (s *Storage) GetPort() string {
	return s.Port
}

func (s *Storage) GetUser() string {
	return s.User
}
func (s *Storage) GetPassword() string {
	return s.Pass
}

func (s *Storage) GetDBName() string {
	return s.DbName
}
