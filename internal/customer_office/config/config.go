package config

type Config struct {
	Port     int `env:"SERVER_PORT" envDefault:"5000"`
	HttpPort int `env:"HTTP_GATEWAY_PORT" envDefault:"5100"`

	RestaurantMenuHost string `env:"RESTAURANT_MENU_HOST" envDefault:"host.docker.internal"`
	RestaurantMenuPort int    `env:"RESTAURANT_MENU_PORT" envDefault:"5002"`

	AMQPHost     string `env:"AMQP_HOST" envDefault:"host.docker.internal"`
	AMQPPort     int    `env:"AMQP_PORT" envDefault:"5201"`
	AMQPUsername string `env:"AMQP_USERNAME" envDefault:"guest"`
	AMQPPassword string `env:"AMQP_PASSWORD" envDefault:"guest"`

	AMQPOrderRestaurantQueue string `env:"AMQP_ORDER_RESTAURANT_QUEUE" envDefault:"orders-restaurant"`
	AMQPOrderStatisticsQueue string `env:"AMQP_ORDER_STATISTICS_QUEUE" envDefault:"orders-statistics"`

	DBPort     int    `env:"DB_PORT" envDefault:"5101"`
	DBHost     string `env:"DB_HOST" envDefault:"host.docker.internal"`
	DBName     string `env:"DB_NAME" envDefault:"db"`
	DBUsername string `env:"DB_USER" envDefault:"db"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"db"`
}
