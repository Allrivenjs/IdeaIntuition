package ChatHistory

import (
	User "IdeaIntuition/app/models/User"
	"IdeaIntuition/global"
	"IdeaIntuition/services"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name        string        `gorm:"size:100;not null" json:"name"`
	Description string        `gorm:"size:255" json:"description"`
	User        User.User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	ChatHistory []ChatHistory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Reason      Reason        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type ChatHistory struct {
	gorm.Model
	Message string    `json:"message"`
	User    User.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Reason struct {
	gorm.Model
	services.PromptListProjectStruct
}

func (r *Room) Create() {
	if err := global.DB.Create(&r).Error; err != nil {
		panic(err)
	}
}

func (r *Room) GetChatHistoryByUser(u User.User) []ChatHistory {
	var chatHistory []ChatHistory
	global.DB.Model(&r).
		Where("user_id = ?", u.ID).
		Preload("User").
		Preload("ChatHistory.User").
		Preload("Reason").
		Find(&chatHistory)
	return chatHistory
}
