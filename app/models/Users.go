package models

import (
	"IdeaIntuition/global"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string         `gorm:"size:100;not null" json:"first_name"`
	LastName  string         `gorm:"size:100;not null" json:"last_name"`
	Email     string         `gorm:"size:100;unique;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"-"`
	Interests []UserInterest `json:"interests"`
	Active    bool           `gorm:"default:true" json:"active"`
}

type Interest struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
}

type UserInterest struct {
	UserID     uint     `gorm:"primaryKey" json:"user_id"`
	InterestID uint     `gorm:"primaryKey" json:"interest_id"`
	Score      float64  `gorm:"type:decimal(5,2);not null" json:"score"`
	User       User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Interest   Interest `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := global.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
