package config

import "time"

type Config interface {
	GetApp() string
	GetHTTPServer() HTTPServer
	GetStorage() Storage
}

type HTTPServer interface {
	GetAddress() string
	GetTimeout() time.Duration
	GetIdleTimeout() time.Duration
}

type Storage interface {
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetDBName() string
}
