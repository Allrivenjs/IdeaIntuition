package models

import (
	"IdeaIntuition/global"
	"IdeaIntuition/services"
	"errors"
	"fmt"
)

type Room struct {
	Model
	Name        string        `gorm:"size:100;not null" json:"name"`
	Description string        `gorm:"size:255" json:"description"`
	UserId      uint          `gorm:"not null" json:"-"`
	ChatHistory []ChatHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RoomId" json:"-"`
	ReasonId    uint          `gorm:"not null" json:"-"`
	Reason      Reason        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ReasonId;references:ID" json:"-"`
}

type ChatHistory struct {
	Model
	Message string `json:"message"`
	UserId  uint   `gorm:"not null" json:"-"`
	RoomId  uint   `gorm:"not null" json:"-"`
}

type Reason struct {
	Model
	services.PromptListProjectStruct
}

func (r *Room) Create() {
	if err := global.DB.Create(&r).Error; err != nil {
		panic(err)
	}
}

func (r *Room) GetChatHistoryByUser(u User) []ChatHistory {
	var chatHistory []ChatHistory
	global.DB.Model(&r).
		Where("user_id = ?", u.ID).
		Preload("User").
		Preload("ChatHistory.User").
		Preload("Reason").
		Find(&chatHistory)
	return chatHistory
}

func (r *Room) Load(param interface{}) error {
	switch p := param.(type) {
	case string:
		fmt.Printf("El parámetro es un string: %s\n", p)
	case []string:
		fmt.Printf("El parámetro es un array: %v\n", p)
	default:
		return errors.New("param is not a string or an array")
	}
	return nil
}
