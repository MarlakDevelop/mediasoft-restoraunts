package orderproduceramqp

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	customer "restaurant/pkg/contracts/customer_office"
)

type OrderProducerAmqp struct {
	ch     *amqp.Channel
	queues []*amqp.Queue
}

func New(ch *amqp.Channel, qNames []string) *OrderProducerAmqp {
	var queues []*amqp.Queue
	for _, qName := range qNames {
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
		queues = append(queues, &queue)
	}
	return &OrderProducerAmqp{
		ch:     ch,
		queues: queues,
	}
}

func (p *OrderProducerAmqp) SendOrder(order *customer.OrderJSON) {
	for _, queue := range p.queues {
		data, err := json.Marshal(order)
		if err != nil {
			return
		}
		msg := amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		}
		err = p.ch.Publish(
			"",
			queue.Name,
			false,
			false,
			msg,
		)
		if err != nil {
			log.Errorf("Sending order to queue(%s) error - %s", queue.Name, err.Error())
			return
		}
	}
}
