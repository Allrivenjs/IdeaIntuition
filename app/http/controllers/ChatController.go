package controllers

import (
	"IdeaIntuition/app/models"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type createProjectRequest struct {
	models.Reason
	User_id uint `json:"user_id" valid:"required"`
}

func CreateProject(c *fiber.Ctx) error {
	var body createProjectRequest
	if err := validateProjectParams(c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	return nil
}

func validateProjectParams(c *fiber.Ctx) error {
	validateStruct, err := govalidator.ValidateStruct(&createProjectRequest{})
	if validateStruct {
		return err
	}
	return nil
}
