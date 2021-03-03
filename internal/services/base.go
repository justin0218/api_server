package services

import "api_server/store"

type baseService struct {
	Config   store.Config
	Rabbitmq store.Rabbitmq
}
