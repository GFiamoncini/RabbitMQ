package main

import (
	"github.com/GFiamoncini/RabbitMQ/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Teste do GF", "amq.direct")
}
