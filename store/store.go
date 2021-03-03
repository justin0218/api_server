package store

import (
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
	"sync"
)

var (
	logOnce      sync.Once
	rabbitMqOnce sync.Once
	configOnce   sync.Once
)

var (
	logClient      *logs.BeeLogger
	rabbitMqClient *amqp.Connection
	config         cfg
)
