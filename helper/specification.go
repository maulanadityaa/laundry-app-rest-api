package helper

import (
	"strings"

	"gorm.io/gorm"
)

func SelectByName(customerName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(customerName)+"%")
	}
}
