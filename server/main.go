package main

import (
	"fmt"
	"github.com/Niladri2003/server-monitor/server/consume"
	"github.com/Niladri2003/server-monitor/server/pkg/configs"
	"github.com/Niladri2003/server-monitor/server/pkg/middleware"
	"github.com/Niladri2003/server-monitor/server/pkg/routes"
	"github.com/Niladri2003/server-monitor/server/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config := configs.FiberConfig()

	//Define a new fiber app with config.
	app := fiber.New(config)

	//Middlewares
	middleware.FiberMiddleware(app)
	influxconfig := utils.NewConfig()

	go func() {
		if err := consume.ConsumeKafka(influxconfig.InfluxClient); err != nil {
			log.Fatalf("Failed to consume Kafka messages: %v", err)
		}
	}()

	// Pass the InfluxClient to the route handlers
	routes.PublicRoutes(app, influxconfig)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start server (with or without graceful shutdown)
	if os.Getenv("STAGE_STATUS") == "dev" {
		fmt.Println("Running in development mode")
		fmt.Println(os.Getenv("STAGE_STATUS"))
		utils.StartServer(app)
	} else {
		fmt.Println("Running in production mode")
		//utils.StartServerWithGracefulShutdown(app)
		utils.StartServer(app)
	}
}
