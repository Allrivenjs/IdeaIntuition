package models

import (
	"IdeaIntuition/global"
	"gorm.io/gorm"
)

type Idea struct {
	gorm.Model
	Content  string `gorm:"size:500;not null" json:"content"`
	RoomId   uint   `gorm:"not null" json:"-"`
	Room     Room   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:RoomId;references:ID" json:"-"`
	Selected bool   `gorm:"not null" json:"selected"`
}

func (i *Idea) Create() {
	if err := global.DB.Create(&i).Error; err != nil {
		panic(err)
	}
}

func (i *Idea) GetIdeaById(id uint) error {
	if err := global.DB.First(&i, id).Error; err != nil {
		return err
	}
	return nil
}

func (i *Idea) Update() {
	if err := global.DB.Save(&i).Error; err != nil {
		panic(err)
	}
}

func GetIdeas(roomID uint) ([]Idea, error) {
	var ideas []Idea
	if err := global.DB.Where("room_id = ?", roomID).Find(&ideas).Error; err != nil {
		return nil, err
	}
	return ideas, nil
}
