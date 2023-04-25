package main

import (
	c "IdeaIntuition/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	logrus.Infof("beginning application")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	// Inicializa la base de datos
	c.InitDB()
}
