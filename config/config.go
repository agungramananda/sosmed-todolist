package config

import (
	"os"
)

var conf Config

type Config struct {
	ServiceName    	string
	ServiceHost		string
	ServicePort		string
	SwaggerHost    	string
	SwaggerPort		string
	DbConf         	*DBConfig
}

func New() *Config {
	conf = Config{
		ServiceName:		os.Getenv("SVC_NAME"),
		ServiceHost: 		os.Getenv("SVC_HOST"),
		ServicePort: 		os.Getenv("SVC_PORT"),
		SwaggerHost: 		os.Getenv("SWAGGER_HOST"),
		SwaggerPort: 		os.Getenv("SWAGGER_PORT"),
		DbConf: 			&DBConfig{
			Host: 		os.Getenv("DB_HOST"),
			Port: 		os.Getenv("DB_PORT"),
			Username: 	os.Getenv("DB_USERNAME"),
			Password: 	os.Getenv("DB_PASSWORD"),
			Database: 	os.Getenv("DB_NAME"),
		},
	}

	return &conf
}

func Get() *Config {
	return &conf
}