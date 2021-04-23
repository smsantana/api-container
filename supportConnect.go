package main

import (
	"os"
)

// RedisConfig config redis
type RedisConfig struct {
	redis_host     string
	redis_port     string
	redis_password string
}

// RedisConfig config redis
type DbConfig struct {
	dbdriver string
	host     string
	password string
	user     string
	dbname   string
	port     string
}

// DbConnect connect
func DbConnect() *DbConfig {
	result := DbConfig{os.Getenv("DB_DRIVER"), os.Getenv("DB_HOST"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT")}
	return &result
}

// RedisConnect teste ok
func RedisConnect() *RedisConfig {
	//redis details
	resultado := RedisConfig{os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASSWORD")}
	return &resultado
}
