package main

import (
	"api_server/internal/routers"
	"api_server/store"
	"fmt"
	"time"
)

func init() {
	log := new(store.Log)
	log.Get().Debug("server started at %v", time.Now())
}

func main() {
	mq := new(store.Rabbitmq)
	mq.Get()
	config := new(store.Config)
	err := routers.Init().Run(fmt.Sprintf(":%d", config.Get().Http.Port))
	if err != nil {
		panic(err)
	}
}
