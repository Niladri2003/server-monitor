package controllers

import (
	"fmt"
	"github.com/Niladri2003/server-monitor/server/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	userId, err := middleware.GetUserId(c)
	fmt.Println(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	id, err := middleware.ExtractTokenMetadata(c)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"userId": userId,
		"id":     id,
	})
}
