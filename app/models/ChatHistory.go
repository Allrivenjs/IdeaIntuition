package models

import (
	"IdeaIntuition/services"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name        string        `gorm:"size:100;not null" json:"name"`
	Description string        `gorm:"size:255;not null" json:"description"`
	User        User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	ChatHistory []ChatHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type ChatHistory struct {
	gorm.Model
	Message string `json:"message"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Reason struct {
	gorm.Model
	services.PromptListProjectStruct
}

func (r *Room) create() {

}
