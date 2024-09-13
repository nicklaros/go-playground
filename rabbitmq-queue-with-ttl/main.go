package main

import (
	"log"

	"github.com/streadway/amqp"
)

func mains() {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://user:user@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue with TTL (in milliseconds)
	args := make(amqp.Table)
	args["x-message-ttl"] = 1000 * 60 * 60 * 24 * 7 // 60,000 milliseconds = 1 minute
	args["x-queue-type"] = "quorum"

	_, err = ch.QueueDeclare(
		"queue-with-message-ttl",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		args,  // arguments with x-message-ttl
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
}
