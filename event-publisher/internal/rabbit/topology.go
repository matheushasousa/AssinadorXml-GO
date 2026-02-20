package rabbit

import "github.com/rabbitmq/amqp091-go"

func InitTopology(conn *amqp091.Connection, cfg Config) {

	ch, _ := conn.Channel()

	ch.ExchangeDeclare("xml.events", "topic", true, false, false, false, nil)

	ch.QueueDeclare("xml.recebido.queue", true, false, false, false, nil)
	ch.QueueBind("xml.recebido.queue", "xml.recebido", "xml.events", false, nil)

	ch.QueueDeclare("xml.assinado.queue", true, false, false, false, nil)
	ch.QueueBind("xml.assinado.queue", "xml.assinado", "xml.events", false, nil)
}
