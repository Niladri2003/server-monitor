package routes

import (
	"github.com/Niladri2003/server-monitor/server/app/controllers"
	"github.com/Niladri2003/server-monitor/server/pkg/middleware"
	utils2 "github.com/Niladri2003/server-monitor/server/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App, config *utils2.InfluxConfig) {
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
	route.Get("/server/memory-usage", func(c *fiber.Ctx) error {
		return controllers.GetMemoryUsage(c, config.InfluxClient)
	})
	route.Get("/server/swap-usage", func(c *fiber.Ctx) error {
		return controllers.GetSwapMemoryUsage(c, config.InfluxClient)
	})
	route.Get("/server/cpu-usage", func(c *fiber.Ctx) error {
		return controllers.GetCpuUsage(c, config.InfluxClient)
	})
	route.Get("/server/get-top5-process-by-cpu", func(c *fiber.Ctx) error {
		return controllers.GetTop5ProcessByCpu(c, config.InfluxClient)
	})
	route.Get("/server/get-top5-process-by-memory", func(c *fiber.Ctx) error {
		return controllers.GetTop5ProcessByMemory(c, config.InfluxClient)
	})
	route.Get("/server/get-server-info", func(c *fiber.Ctx) error {
		return controllers.GetHostInfo(c, config.InfluxClient)
	})
	route.Post("/test", middleware.Protected(), controllers.Test)
	route.Post("/add-server", middleware.Protected(), controllers.AddServerForMonitoring)
	route.Put("/update-server/:id", middleware.Protected(), controllers.UpdateServer)
	route.Delete("/delete-server/:id", middleware.Protected(), controllers.DeleteServer)
	route.Get("/get-all-servers", middleware.Protected(), controllers.GetServersByUser)
	route.Get("/get-server/:id", middleware.Protected(), controllers.GetServerByID)

}
