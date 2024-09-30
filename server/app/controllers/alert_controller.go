package controllers

import (
	"context"
	"github.com/Niladri2003/server-monitor/server/app/models"
	"github.com/Niladri2003/server-monitor/server/pkg/middleware"
	"github.com/Niladri2003/server-monitor/server/pkg/utils"
	"github.com/Niladri2003/server-monitor/server/platform/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateAlert(c *fiber.Ctx) error {
	var alert models.Alert
	now := time.Now().Unix()
	claims, err := middleware.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "token invalid",
			"status":  fiber.StatusInternalServerError,
		})
	}
	expires := claims.Expires
	//Checking if now time is greater than expiration from jwt
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   true,
			"message": "token expired",
			"status":  fiber.StatusUnauthorized,
		})
	}
	if err := c.BodyParser(&alert); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "parse body failed",
		})
	}
	serverID := c.Params("id")
	if serverID == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "server id is empty",
		})
	}
	alert.Server, err = primitive.ObjectIDFromHex(serverID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "parse server id failed",
		})
	}
	// Validate Alert fields
	if err := validate.Struct(alert); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := utils.ValidatorErrors(validationErrors) // Format errors if needed
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   errorMessages,
		})
	}

	alert.Status = "active"
	alert.CreatedAt = time.Now()

	//insert Alert into the database
	collection := database.GetDbCollection("alert")
	result, err := collection.InsertOne(context.TODO(), alert)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "create alert failed",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "create alert success",
		"data":  result,
	})

}
