package main

import (
	configpb "github.com/toaddyan/redis-slinker/pkg/config/pb"
)

type Config struct {
	server  configpb.Server
	options configpb.Options
	redis   configpb.Redis
}

func GetConfig() *Config {
	return &Config{
		server: configpb.Server{
			Port: "8080",
		},
		options: configpb.Options{
			Schema: "http",
			Prefix: "localhost:8080",
		},
		redis: configpb.Redis{
			Host: "redis",
			Port: "6379",
		},
	}
}
