package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/cauakath/timeline-server/domain"
	"github.com/cauakath/timeline-server/enum"
	"github.com/cauakath/timeline-server/model"
	"github.com/gofiber/fiber/v2"
)

type TimelineController struct {
	timelineUseCase domain.TimelineUseCase
}

func NewTimelineController(timelineUseCase domain.TimelineUseCase) *TimelineController {
	return &TimelineController{
		timelineUseCase: timelineUseCase,
	}
}

func (n *TimelineController) CreateTimeline(ctx *fiber.Ctx) error {
	var timelineRequest model.Timeline
	var response model.Response

	if err := ctx.BodyParser(&timelineRequest); err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := VerifyRequest(timelineRequest, ctx); err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := n.timelineUseCase.CreateTimeline(timelineRequest); err != nil {
		response = model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		StatusCode: http.StatusCreated,
		Message:    "success",
	}

	return ctx.Status(http.StatusCreated).JSON(response)
}

func (n *TimelineController) UpdateTimeline(ctx *fiber.Ctx) error {
	var timelineRequest model.Timeline
	var response model.Response

	id := ctx.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if err := ctx.BodyParser(&timelineRequest); err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	timeline, err := n.timelineUseCase.GetTimeline(idInt)
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if timeline.Title == "" {
		response = model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "notFound",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	if err := n.timelineUseCase.UpdateTimeline(timelineRequest, idInt); err != nil {
		response = model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		StatusCode: http.StatusCreated,
		Message:    "success",
	}

	return ctx.Status(http.StatusCreated).JSON(response)
}

func VerifyRequest(timelineRequest model.Timeline, ctx *fiber.Ctx) error {
	if timelineRequest.Title == "" {
		return errors.New("title is required")
	}

	if timelineRequest.Type == "" {
		return errors.New("type is required")
	}

	if timelineRequest.Type != enum.Work && timelineRequest.Type != enum.Education && timelineRequest.Type != enum.Other {
		return errors.New("type must be WORK, EDUCATION, or OTHER")
	}

	if timelineRequest.Location == "" {
		return errors.New("location is required")
	}

	if timelineRequest.Start == "" {
		return errors.New("start is required")
	}

	return nil
}

func (n *TimelineController) GetTimeline(ctx *fiber.Ctx) error {
	var response model.Response
	var id = ctx.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	timeline, err := n.timelineUseCase.GetTimeline(idInt)
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	if timeline.Title == "" {
		response = model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "notFound",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	response = model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       timeline,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (n *TimelineController) ListTimelines(ctx *fiber.Ctx) error {
	var response model.Response

	timelines, err := n.timelineUseCase.ListTimelines()
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       timelines,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

func (n *TimelineController) DeleteTimeline(ctx *fiber.Ctx) error {
	var response model.Response
	var id = ctx.Params("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	timeline, err := n.timelineUseCase.GetTimeline(idInt)
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if timeline.Title == "" {
		response = model.Response{
			StatusCode: http.StatusNotFound,
			Message:    "notFound",
		}

		return ctx.Status(http.StatusNotFound).JSON(response)
	}

	err = n.timelineUseCase.DeleteTimeline(idInt)
	if err != nil {
		response = model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		return ctx.Status(http.StatusInternalServerError).JSON(response)
	}

	response = model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
	}

	return ctx.Status(http.StatusOK).JSON(response)
}
