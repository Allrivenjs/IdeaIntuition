package config

import (
	m "IdeaIntuition/app/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitDB() {
	// Configura tus credenciales de la base de datos aquí
	logrus.Infof("loading database config")
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	logrus.Infof("%v", dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
		logrus.Errorf("failed to connect database: %v", err)
	}
	//ejecuta las migraciones
	migrate()
}

func migrate() {
	// Reemplaza `User`, `Interest` y `UserInterest` con las estructuras de tus modelos
	logrus.Infof("migrate database")
	err := DB.AutoMigrate(&m.User{}, &m.Interest{}, &m.UserInterest{})
	if err != nil {
		panic("failed to migrate database")
		logrus.Errorf("failed to migrate database: %v", err)
	}
}