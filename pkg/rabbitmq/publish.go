package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func Publish() {

}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	defer conn.Close()

	ch, err := conn.Channel()

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	log.Printf(" [x] Sent %s", body)
}
