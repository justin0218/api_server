package wechat_server

import (
	"api_server/pkg/etcd"
	"sync"
)

var once sync.Once
var conn WechatClient

func GetClient() WechatClient {
	once.Do(func() {
		conn = NewWechatClient(etcd.Discovery("wechat_server"))
	})
	return conn
}
