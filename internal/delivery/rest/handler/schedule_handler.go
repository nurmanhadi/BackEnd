package handler

import (
	"liva/internal/model"
	"liva/internal/service"
	"liva/pkg/exception"

	"github.com/gofiber/fiber/v2"
)

type ScheduleHandler interface {
	AddSchedule(c *fiber.Ctx) error
	UpdateSchedule(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
}
type scheduleHandler struct {
	scheduleService service.ScheduleService
}

func NewScheduleHandler(scheduleService service.ScheduleService) ScheduleHandler {
	return &scheduleHandler{scheduleService: scheduleService}
}

func (h *scheduleHandler) AddSchedule(c *fiber.Ctx) error {
	request := new(model.ScheduleAddRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.scheduleService.AddSchedule(*request); err != nil {
		return err
	}
	return NewResponse(c, 201, "OK")
}
func (h *scheduleHandler) UpdateSchedule(c *fiber.Ctx) error {
	scheduleId := c.Params("scheduleId")
	request := new(model.ScheduleUpdateRequest)
	if err := c.BodyParser(&request); err != nil {
		return exception.NewError(400, "failed parse to json")
	}
	if err := h.scheduleService.UpdateSchedule(scheduleId, *request); err != nil {
		return err
	}
	return NewResponse(c, 200, "OK")
}
func (h *scheduleHandler) FindAll(c *fiber.Ctx) error {
	schedules, err := h.scheduleService.FindAll()
	if err != nil {
		return err
	}
	return NewResponse(c, 200, schedules)
}
func (h *scheduleHandler) FindById(c *fiber.Ctx) error {
	scheduleId := c.Params("scheduleId")
	schedule, err := h.scheduleService.FindById(scheduleId)
	if err != nil {
		return err
	}
	return NewResponse(c, 200, schedule)
}
