package user_server

import (
	"api_server/pkg/etcd"
	"sync"
)

var once sync.Once
var conn UserClient

func GetClient() UserClient {
	once.Do(func() {
		conn = NewUserClient(etcd.Discovery("user_server"))
	})
	return conn
}
