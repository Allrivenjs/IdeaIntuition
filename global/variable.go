package global

import "gorm.io/gorm"

var DB *gorm.DB

func Load(db *gorm.DB) {
	DB = db
}
