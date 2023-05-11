package models

import (
	"IdeaIntuition/global"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
}

func (m *Model) loadRelationsModels(relation string) error {
	if err := global.DB.Preload(relation).Find(&m).Error; err != nil {
		logrus.Errorf("failed to load relations: %v", err)
		return err
	}
	return nil
}
