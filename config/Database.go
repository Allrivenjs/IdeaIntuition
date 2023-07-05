package config

import (
	CH "IdeaIntuition/app/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDB() *gorm.DB {
	// Configura tus credenciales de la base de datos aquí
	logrus.Infof("loading database config")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	logrus.Infof("%v", dsn)
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
		logrus.Errorf("failed to connect database: %v", err)
	}
	//ejecuta las migraciones
	logrus.Infof("migrate database")
	migrate(DB)
	logrus.Infof("seeder on database")
	seedUsers(DB)
	return DB
}

func migrate(DB *gorm.DB) {
	// Reemplaza `U`, `Interest` y `UserInterest` con las estructuras de tus modelos
	err := DB.AutoMigrate(&CH.User{},
		&CH.Interest{},
		&CH.UserInterest{},
		&CH.Reason{},
		&CH.ChatHistory{},
		&CH.Room{})
	if err != nil {
		panic("failed to migrate database")
		logrus.Errorf("failed to migrate database: %v", err)
	}
}
