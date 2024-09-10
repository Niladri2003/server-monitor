package routes

import (
	"github.com/Niladri2003/server-monitor/server/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/user/sign-up", func(c *fiber.Ctx) error {
		return controllers.UserSignUp(c)
	})
	route.Post("/user/sign-in", func(c *fiber.Ctx) error {
		return controllers.UserSignIn(c)
	})

}
