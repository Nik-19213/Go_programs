package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	// Consume messages from the queue
	msgs, err := ch.Consume(
		"test_queue", // queue name
		"",           // consumer name
		true,         // auto-acknowledge
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	log.Println("Waiting for messages...")

	for msg := range msgs {
		log.Printf("Received: %s", msg.Body)
	}
}
