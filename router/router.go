package router

import (
	"github.com/cauakath/timeline-server/controller"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, timelineController *controller.TimelineController) *fiber.App {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello, World!")
	})

	router.Post("/timeline", timelineController.CreateTimeline)
	router.Put("/timeline/:id", timelineController.UpdateTimeline)
	router.Get("/timeline", timelineController.ListTimelines)
	router.Get("/timeline/:id", timelineController.GetTimeline)
	router.Delete("/timeline/:id", timelineController.DeleteTimeline)

	return router
}
