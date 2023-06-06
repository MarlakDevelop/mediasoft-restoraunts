package bootstrap

import (
	"fmt"
	"github.com/streadway/amqp"
	"restaurant/internal/statistics_statistics/config"
)

func InitAMQPConnection(cfg config.Config) (*amqp.Connection, error) {
	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.AMQPUsername, cfg.AMQPPassword, cfg.AMQPHost, cfg.AMQPPort)
	return amqp.Dial(amqpURI)
}

func InitAMQPChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	return conn.Channel()
}

func InitAMQPQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
}
