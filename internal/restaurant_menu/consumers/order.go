package consumers

import (
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type OrderConsumerAmqp struct {
	ch       *amqp.Channel
	queue    *amqp.Queue
	delivery <-chan amqp.Delivery
}

func NewOrderConsumerAmqp(ch *amqp.Channel, qName string) *OrderConsumerAmqp {
	queue, err := ch.QueueDeclare(
		qName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare queue %s: %v", qName, err)
	}

	delivery, err := ch.Consume(
		qName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	consumer := &OrderConsumerAmqp{
		ch:       ch,
		queue:    &queue,
		delivery: delivery,
	}

	return consumer
}

func (c *OrderConsumerAmqp) Handle() {
	// Handle incoming messages
	for d := range c.delivery {
		log.Printf("Received message: %s", d.Body)
	}
}
