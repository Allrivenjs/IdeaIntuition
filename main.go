package main

import (
	"IdeaIntuition/services"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	//logrus.Infof("beginning application")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	//// Inicializa la base de datos
	//global.Load(c.InitDB())
	//app := fiber.New()
	//c.SetupRoutes(app)
	////app.Use(middlewares.RouteLogger(app))
	//logrus.Fatal(app.Listen(":3000"))

	res, err := services.SendMessage([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "¿Puedes proporcionarme una lista de posibles proyectos?",
		},
	}, 300)

	if err != nil {
		log.Fatal(err)
	}

	// Se procesan las opciones de respuesta generadas por el modelo de lenguaje
	logrus.Printf("Respuesta: %v", res)
}
