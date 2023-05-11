package config

import (
	"IdeaIntuition/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func seedUsers(db *gorm.DB) {
	pass, erro := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if erro != nil {
	}
	users := []models.User{
		{
			FirstName: "Admin",
			Email:     "admin@gmail.com",
			Active:    true,
			Password:  string(pass),
			LastName:  "Demon",
		},
	}

	for _, user := range users {
		db.Create(&user)
	}
}
