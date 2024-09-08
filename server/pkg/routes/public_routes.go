package routes

import (
	"github.com/Niladri2003/server-monitor/server/app/controllers"
	utils2 "github.com/Niladri2003/server-monitor/server/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App, config *utils2.InfluxConfig) {
	route := a.Group("/api/v1")

	// Pass the InfluxClient to the controller
	route.Get("/server/metrics", func(c *fiber.Ctx) error {
		return controllers.GetMetrics(c, config.InfluxClient)
	})
	route.Get("/server/disk-usage", func(c *fiber.Ctx) error {
		return controllers.GetDiskUsage(c, config.InfluxClient)
	})
	route.Get("/server/network-stats", func(c *fiber.Ctx) error {
		return controllers.GetNetworkStats(c, config.InfluxClient)
	})
	
}
