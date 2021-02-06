package file_server

import (
	"api_server/pkg/etcd"
	"sync"
)

var once sync.Once
var conn FileClient

func GetClient() FileClient {
	once.Do(func() {
		conn = NewFileClient(etcd.Discovery("file_server"))
	})
	return conn
}
