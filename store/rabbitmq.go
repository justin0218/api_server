package store

import (
	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	conf Config
}

func (s *Rabbitmq) Get() *amqp.Connection {
	rabbitMqOnce.Do(func() {
		var (
			err error
		)
		rabbitMqClient, err = amqp.Dial(s.conf.Get().Rabbitmq.Addr)
		if err != nil {
			panic(err)
		}
	})
	return rabbitMqClient
}
