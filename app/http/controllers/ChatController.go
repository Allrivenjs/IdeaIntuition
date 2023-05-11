package controllers

import (
	"IdeaIntuition/app/models/ChatHistory"
	"IdeaIntuition/app/models/User"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type createProjectRequest struct {
	ChatHistory.Reason
	UserId uint `json:"user_id" valid:"required"`
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
	u, err := User.Find(int(body.UserId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	room := ChatHistory.Room{
		Name:        `Project:` + body.TypeProject,
		Description: "",
		User:        u,
		Reason:      body.Reason,
	}
	room.Create()
	return c.JSON(fiber.Map{
		"room": room,
		"msg":  "Room created successfully",
	})
}

func validateProjectParams(c *fiber.Ctx) error {
	validateStruct, err := govalidator.ValidateStruct(&createProjectRequest{})
	if validateStruct {
		return err
	}
	return nil
}
