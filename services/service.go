package services

import (
	"aneka-zoo/database"
	"log"

	"gorm.io/gorm"
)

func Create(entity interface{}) (result *gorm.DB) {
	result = database.GetConnection().Create(entity)
	if result.Error != nil {
		log.Println("Error created data >>>", result.Error)
	}
	return
}
