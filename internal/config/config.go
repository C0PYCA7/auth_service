package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string   `yaml:"env"`
	DbConfig   Database `yaml:"database"`
	GRPCconfig GRPC     `yaml:"grpc"`
	JWTConfig  JWT      `yaml:"jwt"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"name"`
}

type GRPC struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type JWT struct {
	Key string `yaml:"key"`
}

func LoadConfig() *Config {
	var cfg Config

	path := "././config/config.yml"

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		log.Println("Error reading config:", err)
		os.Exit(1)
	}
	return &cfg
}
