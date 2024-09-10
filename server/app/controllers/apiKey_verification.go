package controllers

import (
	"context"
	"github.com/Niladri2003/server-monitor/server/app/models"
	"github.com/Niladri2003/server-monitor/server/platform/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func ApiKeyVerification(c *fiber.Ctx) error {
	apikey := c.Query("api_key")
	if apikey == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	server := models.Server{}
	collection := database.GetDbCollection("server")
	err := collection.FindOne(context.Background(), bson.M{"api_key": apikey}).Decode(&server)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
