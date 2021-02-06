package services

import "api_server/store"

type baseService struct {
	Redis  store.Redis
	Config store.Config
}
