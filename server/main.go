package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	app := fiber.New()

	go consumeKafka()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}

func consumeKafka() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "agent-data-topic",
		GroupID: "central-server-group",
	})

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		log.Printf("Received data: %s", string(m.Value))
	}
}
