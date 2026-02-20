package rabbit

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func NewConnection(cfg Config) *amqp091.Connection {

	conn, err := amqp091.Dial(cfg.Url)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
