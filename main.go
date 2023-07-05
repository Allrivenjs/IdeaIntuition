package main

import (
	c "IdeaIntuition/config"
	"IdeaIntuition/global"
	"IdeaIntuition/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
	"os"
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
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	routes.SetupRoutes(app)
	//app.Use(middlewares.RouteLogger(app))
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	logrus.Fatal(app.Listen("0.0.0.0:" + port))
}
