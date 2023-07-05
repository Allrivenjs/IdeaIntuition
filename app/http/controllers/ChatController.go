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
	//fmt.Println(u, u.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	reason := models.Reason{
		PromptListProjectStruct: services.PromptListProjectStruct{
			TypeProject:  body.TypeProject,
			Approach:     body.Approach,
			Requirements: body.Requirements,
			Course:       body.Course,
			Technology:   body.Technology,
			Context:      body.Context,
		},
	}
	reason.Create()
	room := models.Room{
		Name:        `Project:` + body.TypeProject,
		Description: "",
		UserId:      u.ID,
		Reason:      reason,
	}
	room.Create()
	room.User = u
	room.Reason = reason
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

type createMessage struct {
	RoomID uint `json:"room_id" valid:"required"`
}

func validateMessagesBody(body createMessage) error {
	v, err := govalidator.ValidateStruct(body)
	logrus.Printf("Validation error: %v, error: %v", v, err)
	if err != nil {
		return err
	}
	return nil
}

func Messages(c *fiber.Ctx) error {

	var body createMessage
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := validateMessagesBody(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	room, err := models.GetRoom(body.RoomID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Room not found",
		})
	}
	err = room.Load([]string{"User", "Reason"})
	if err != nil {
		return err
	}

	p := room.Reason.PromptListProjectStruct
	project, err := p.GetListOfPossibleProject([]openai.ChatCompletionMessage{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}
	content := services.ConfigureMessage(project.Choices[0].Message.Content)

	// Crear un canal para recibir los modelos creados
	modelsCreated := make(chan models.Idea)

	go func() {
		for _, c := range content {
			idea := models.Idea{
				Content:  c,
				RoomId:   room.ID,
				Selected: false,
			}
			idea.Create()

			// Enviar el modelo creado al canal
			modelsCreated <- idea
		}

		// Cerrar el canal después de que se hayan enviado todos los modelos
		close(modelsCreated)
	}()

	// Crear un slice para almacenar los modelos creados
	var createdModels []models.Idea

	// Recorrer el canal para recibir los modelos y agregarlos al slice
	for model := range modelsCreated {
		createdModels = append(createdModels, model)
	}

	return c.JSON(fiber.Map{
		"message":          createdModels,
		"token_completion": project.Usage.CompletionTokens,
		"token_total":      project.Usage.TotalTokens,
		"token_prompt":     project.Usage.PromptTokens,
		"msg":              "Response successfully",
	})
}
