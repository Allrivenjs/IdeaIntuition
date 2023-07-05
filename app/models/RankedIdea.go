package models

import (
	"IdeaIntuition/global"
	"gorm.io/gorm"
)

type RankedIdea struct {
	gorm.Model
	IdeaID uint `gorm:"not null" json:"idea_id"`
	Rank   uint `gorm:"not null" json:"rank"`
	Idea   Idea `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:IdeaID;references:ID" json:"idea"`
}

func (ri *RankedIdea) Create() {
	if err := global.DB.Create(&ri).Error; err != nil {
		panic(err)
	}
}

func (ri *RankedIdea) CreateAndSelectIdea() {
	// Crear la ranked idea
	ri.Create()

	// Buscar la idea correspondiente por su ID
	idea := Idea{}
	if err := idea.GetIdeaById(ri.IdeaID); err != nil {
		panic(err)
	}

	// Marcar la idea como selected: true y asignarle el rank
	idea.Selected = true

	// Actualizar la idea en la base de datos
	if err := global.DB.Save(&idea).Error; err != nil {
		panic(err)
	}
}
