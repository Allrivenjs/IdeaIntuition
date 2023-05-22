package models

import (
	"IdeaIntuition/global"
	"IdeaIntuition/services"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name        string        `gorm:"size:100;not null" json:"name"`
	Description string        `gorm:"size:255" json:"description"`
	UserId      uint          `gorm:"not null" json:"user_id"`
	ChatHistory []ChatHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RoomId" json:"chat_history"`
	ReasonId    uint          `gorm:"not null" json:"reason_id"`
	Reason      Reason        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ReasonId;references:ID" json:"reason"`
	User        User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserId;references:ID" json:"user"`
}

type ChatHistory struct {
	gorm.Model
	Message string `json:"message"`
	UserId  uint   `gorm:"not null" json:"user_id"`
	RoomId  uint   `gorm:"not null" json:"room_id"`
}

type Reason struct {
	gorm.Model
	services.PromptListProjectStruct
}

func (r *Reason) Create() {
	if err := global.DB.Create(&r).Error; err != nil {
		panic(err)
	}
}

func (r *Room) loadRelationsModels(relation string) error {
	// Resto de la implementación
	if err := global.DB.Model(&r).Preload(relation).Find(&r).Error; err != nil {
		logrus.Errorf("failed to load relations: %v", err)
		return err
	}
	return nil
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
		err := r.loadRelationsModels(p)
		if err != nil {
			return err
		}
	case []string:
		for _, relation := range p {
			logrus.Infof("loading relations: %v", relation)
			err := r.loadRelationsModels(relation)
			if err != nil {
				return err
			}
		}
	default:
		return errors.New("param is not a string or an array")
	}
	return nil
}

func GetRoom(id uint) (Room, error) {
	var room Room
	err := global.DB.Find(&room, id).Error
	return room, err
}
