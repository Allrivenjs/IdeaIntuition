package controllers

import (
	"IdeaIntuition/app/models"
	"IdeaIntuition/services"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type createProjectRequest struct {
	models.Reason
	UserId uint `json:"user_id" valid:"required"`
}

func CreateProject(c *fiber.Ctx) error {
	var body createProjectRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := validateProjectParams(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	u, err := models.Find(int(body.UserId))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	room := models.Room{
		Name:        `Project:` + body.TypeProject,
		Description: "",
		UserId:      u.ID,
		Reason:      body.Reason,
	}
	room.Create()
	//load relations of room
	err = room.Load("User")
	if err != nil {
		return err
	}
	err = room.Load([]string{"Reason"})
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"room": room,
		"msg":  "Room created successfully",
	})
}

func validateProjectParams(body createProjectRequest) error {
	v, err := govalidator.ValidateStruct(body)
	logrus.Printf("Validation error: %v, error: %v", v, err)
	if err != nil {
		return err
	}
	return nil
}

func GetMessages(c *fiber.Ctx) error {
	p := services.PromptListProjectStruct{
		TypeProject:  "Creacion de plataforma para el desarrollo de estudios de mercado en programacion",
		Approach:     "temas de educacion",
		Requirements: "tesis",
		Course:       "ingeniera en sistemas.",
		Technology:   "web",
	}
	project, err := p.GetListOfPossibleProject([]openai.ChatCompletionMessage{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": project,
		"msg":     "Response successfully",
	})
}
