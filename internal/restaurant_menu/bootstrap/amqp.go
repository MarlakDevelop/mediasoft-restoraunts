package bootstrap

import (
	"fmt"
	"github.com/streadway/amqp"
	"restaurant/internal/restaurant_menu/config"
)

func InitAMQPConnection(cfg config.Config) (*amqp.Connection, error) {
	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.AMQPUsername, cfg.AMQPPassword, cfg.AMQPHost, cfg.AMQPPort)
	return amqp.Dial(amqpURI)
}
