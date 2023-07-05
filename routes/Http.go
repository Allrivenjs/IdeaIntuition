package routes

import (
	"IdeaIntuition/app/http/controllers"
	"IdeaIntuition/app/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//TODO: Agregar rutas aquí
	api := app.Group("/api")
	setupPublicRoutesApi(api)
	setupProtectedRoutesApi(api)
}

func setupPublicRoutesApi(app fiber.Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hello, World!")
		if err != nil {
			return err
		}
		return nil
	})

	app.Post("/messages", controllers.Messages)

	app.Post("/login", controllers.Login)
	app.Get("/restricted", controllers.Restricted)
	app.Post("/register", controllers.Register)
	app.Post("/create-project", controllers.CreateProject)
	app.Get("/select-idea/:id", controllers.SelectIdea)
	app.Get("/get-ideas/:id", controllers.GetIdeas)
	app.Post("/rank-idea", controllers.RankIdea)
	// Otras rutas públicas
}

func setupProtectedRoutesApi(app fiber.Router) {
	_ = app.Group("/", middlewares.AuthRequired()) // Aplica el middlewares JWT a todas las rutas bajo "/api"
	//Otras rutas protegidas
}
