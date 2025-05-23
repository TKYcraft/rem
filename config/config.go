package config

import "os"

type ServerConfig struct {
	DISCORD_TOKEN  string
	DISCORD_PREFIX string
}

func LoadConfig() (*ServerConfig, error) {
	return &ServerConfig{
		DISCORD_TOKEN:  os.Getenv("DISCORD_BOT_TOKEN"),
		DISCORD_PREFIX: os.Getenv("DISCORD_PREFIX"),
	}, nil
}

type DBConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSLMODE  string
}

func LoadDBConfig() (*DBConfig, error) {
	return &DBConfig{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}, nil
}
