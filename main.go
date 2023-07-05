package main

import (
	c "IdeaIntuition/config"
	"IdeaIntuition/global"
	"IdeaIntuition/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	//logrus.Infof("beginning application")
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf("err loading: %v", err)
	//}
	//// Inicializa la base de datos
	global.Load(c.InitDB())
	app := fiber.New()
	routes.SetupRoutes(app)
	//app.Use(middlewares.RouteLogger(app))
	logrus.Fatal(app.Listen(":3000"))

}
