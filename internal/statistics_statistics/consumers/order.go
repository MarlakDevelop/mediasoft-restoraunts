package consumers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"restaurant/internal/statistics_statistics/repositories/orderrepository"
	"restaurant/internal/statistics_statistics/repositories/productrepository"
	customer "restaurant/pkg/contracts/customer_office"
	"restaurant/pkg/data/slice"
)

type OrderConsumerAmqp struct {
	ch                *amqp.Channel
	queue             *amqp.Queue
	delivery          <-chan amqp.Delivery
	orderRepository   orderrepository.OrderRepository
	productRepository productrepository.ProductRepository
}

func NewOrderConsumerAmqp(ch *amqp.Channel, qName string, orderRepository orderrepository.OrderRepository, productRepository productrepository.ProductRepository) *OrderConsumerAmqp {
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
		ch:                ch,
		queue:             &queue,
		delivery:          delivery,
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}

	return consumer
}

func (c *OrderConsumerAmqp) Handle() {
	// Handle incoming messages
	for d := range c.delivery {
		var order customer.OrderJSON
		err := json.Unmarshal(d.Body, &order)
		if err != nil {
			log.Error(err)
			continue
		}
		products, err := c.productRepository.List(slice.Map[*customer.OrderItemJSON, string](order.Items, func(index int, value *customer.OrderItemJSON, slice []*customer.OrderItemJSON) string {
			return value.ProductUuid
		}))
		if err != nil {
			log.Error(err)
			continue
		}
		err = c.orderRepository.SaveOrderedProducts(slice.Map[*customer.OrderItemJSON, *orderrepository.OrderedProduct](order.Items, func(index int, item *customer.OrderItemJSON, items []*customer.OrderItemJSON) *orderrepository.OrderedProduct {
			productI := slice.IndexOf[*productrepository.Product](products, func(index int, product *productrepository.Product, products []*productrepository.Product) bool {
				return item.ProductUuid == product.Uuid
			})
			var product *productrepository.Product
			if productI != -1 {
				product = products[productI]
			}
			return &orderrepository.OrderedProduct{
				Uuid:        product.Uuid,
				Name:        product.Name,
				Price:       product.Price,
				Count:       int64(item.Count),
				ProductType: int32(product.Type),
				CreatedAt:   order.CreatedAt,
			}
		}))
		if err != nil {
			log.Error(err)
			continue
		}
	}
}
