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
	router.Get("/timeline", timelineController.ListTimelines)

	return router
}
