package main

import (
	RabbitMQ "rabbitmq-miaosha/rabbitmq"
)

func main() {
	rabbitmqConsumeSimple := RabbitMQ.NewRabbitMQSimple("blue")
	rabbitmqConsumeSimple.ConsumeSimple()
}
