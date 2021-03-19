package main

import (
	"fmt"
	"log"

	"github.com/toaddyan/redis-slinker/pkg/handler"
	"github.com/toaddyan/redis-slinker/pkg/service"
	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Println("Starting Slinker")
	configuration := GetConfig()

	pool, err := service.NewPool(configuration.redis.GetHost(), configuration.redis.GetPort(), "")
	if err != nil {
		log.Panic(err)
		return
	}

	service := service.NewService(pool)

	defer service.Close()

	router := handler.New(configuration.options.GetSchema(), configuration.options.GetPrefix(), service)

	log.Fatal(fasthttp.ListenAndServe(":"+configuration.server.Port, router.Handler))
}
