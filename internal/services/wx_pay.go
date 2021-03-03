package services

import (
	"api_server/internal/models/wx"
	"encoding/json"
	"github.com/streadway/amqp"
)

type WxPay struct {
	baseService
}

func (s *WxPay) WxPayNotice(in *wx.WxpayReq) (err error) {
	mq := s.Rabbitmq.Get()
	mqChan, e := mq.Channel()
	if e != nil {
		err = e
		return
	}
	defer mqChan.Close()
	q, e := mqChan.QueueDeclare(
		"wxPayNotice", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if e != nil {
		err = e
		return
	}
	body, e := json.Marshal(in)
	if e != nil {
		err = e
		return
	}
	err = mqChan.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	return
}
