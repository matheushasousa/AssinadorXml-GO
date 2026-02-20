package rabbit

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	ch *amqp091.Channel
}

func NewPublisher(conn *amqp091.Connection, cfg Config) *Publisher {

	ch, _ := conn.Channel()

	ch.ExchangeDeclare(
		"xml.events",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	return &Publisher{ch}
}

func (p *Publisher) Publish(routingKey string, message any) {

	body, _ := json.Marshal(message)

	err := p.ch.Publish(
		"xml.events",
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp091.Persistent,
			Body:         body,
		},
	)

	if err != nil {
		log.Println("Erro publish:", err)
	}
}
