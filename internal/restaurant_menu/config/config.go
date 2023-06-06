package config

type Config struct {
	Port     int `env:"SERVER_PORT" envDefault:"5000"`
	HttpPort int `env:"HTTP_GATEWAY_PORT" envDefault:"5100"`

	CustomerOfficeHost string `env:"CUSTOMER_OFFICE_HOST" envDefault:"host.docker.internal"`
	CustomerOfficePort int    `env:"CUSTOMER_OFFICE_PORT" envDefault:"5001"`

	AMQPHost     string `env:"AMQP_HOST" envDefault:"host.docker.internal"`
	AMQPPort     int    `env:"AMQP_PORT" envDefault:"5201"`
	AMQPUsername string `env:"AMQP_USERNAME" envDefault:"guest"`
	AMQPPassword string `env:"AMQP_PASSWORD" envDefault:"guest"`

	AMQPOrderQueue string `env:"AMQP_ORDER_QUEUE" envDefault:"orders-restaurant"`

	DBPort     int    `env:"DB_PORT" envDefault:"5102"`
	DBHost     string `env:"DB_HOST" envDefault:"host.docker.internal"`
	DBName     string `env:"DB_NAME" envDefault:"db"`
	DBUsername string `env:"DB_USER" envDefault:"db"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"db"`
}
