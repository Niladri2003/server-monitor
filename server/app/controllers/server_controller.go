package controllers

import (
	"context"
	"github.com/Niladri2003/server-monitor/server/app/models"
	"github.com/Niladri2003/server-monitor/server/pkg/middleware"
	"github.com/Niladri2003/server-monitor/server/pkg/utils"
	"github.com/Niladri2003/server-monitor/server/platform/database"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AddServerRequest struct {
	Name          string `json:"name" validate:"required"`
	Description   string `json:"description" validate:"required" `
	CloudProvider string `json:"cloudProvider" bson:"cloud_provider"`
}
type UpdateServerRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	CloudProvider string `json:"cloudProvider"`
}

func AddServerForMonitoring(c *fiber.Ctx) error {
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
	// Parse request body
	req := new(AddServerRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}
	// Validate sign-up fields
	if err := validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := utils.ValidatorErrors(validationErrors) // Format errors if needed
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   errorMessages,
		})
	}

	// Generate UUID for the server ID and API key
	serverID := uuid.New().String()
	apiKey := uuid.New().String()

	userId := claims.UserID

	server := models.Server{
		ID:            primitive.NewObjectID(),
		UserID:        userId,
		ServerID:      serverID,
		APIKey:        apiKey,
		Name:          req.Name,
		Description:   req.Description,
		CloudProvider: req.CloudProvider,
		CreatedAt:     time.Unix(now, 0),
		UpdatedAt:     time.Unix(now, 0),
	}

	// Insert the server into Mongodb
	collection := database.GetDbCollection("server")
	_, err = collection.InsertOne(context.Background(), server)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to add server",
		})

	}
	// Respond with server ID and API key
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Server added successfully",
		"server_id": server.ServerID,
		"api_key":   server.APIKey,
	})
}

func UpdateServer(c *fiber.Ctx) error {
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
	serverID := c.Params("id")

	// Parse request body
	req := new(UpdateServerRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid request body",
		})
	}

	// Prepare update fields
	updateFields := bson.M{}
	if req.Name != "" {
		updateFields["name"] = req.Name
	}
	if req.Description != "" {
		updateFields["description"] = req.Description
	}
	if req.CloudProvider != "" {
		updateFields["cloud_provider"] = req.CloudProvider
	}
	updateFields["updated_at"] = time.Now()

	// Update server details in the database
	collection := database.GetDbCollection("server")
	_, err = collection.UpdateOne(context.Background(), bson.M{"server_id": serverID}, bson.M{"$set": updateFields})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to update server",
		})
	}

	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Server updated successfully",
	})
}
func DeleteServer(c *fiber.Ctx) error {
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
	serverID := c.Params("id")

	// Delete server from the database
	collection := database.GetDbCollection("server")
	_, err = collection.DeleteOne(context.Background(), bson.M{"server_id": serverID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to delete server",
		})
	}

	// Return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Server deleted successfully",
	})
}
func GetServersByUser(c *fiber.Ctx) error {
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

	userId := claims.UserID

	// Find all servers for the user
	var servers []models.Server
	collection := database.GetDbCollection("server")
	cursor, err := collection.Find(context.Background(), bson.M{"user": userId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Failed to retrieve servers",
		})
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &servers); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Error reading servers",
		})
	}

	// Return server list
	return c.Status(fiber.StatusOK).JSON(servers)
}
func GetServerByID(c *fiber.Ctx) error {
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
	serverID := c.Params("id")
	//fmt.Println(serverID)

	// Find server by serverID in the database
	var server models.Server
	collection := database.GetDbCollection("server")
	err = collection.FindOne(context.Background(), bson.M{"server_id": serverID}).Decode(&server)
	if err != nil {
		//fmt.Println("error", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Server not found",
		})
	}

	// Return server details
	return c.Status(fiber.StatusOK).JSON(server)
}
