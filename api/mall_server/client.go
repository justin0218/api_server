package mall_server

import (
	"api_server/pkg/etcd"
	"sync"
)

var once sync.Once
var conn MallClient

func GetClient() MallClient {
	once.Do(func() {
		conn = NewMallClient(etcd.Discovery("mall_server"))
	})
	return conn
}
