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

func Update(entity interface{}) (result *gorm.DB) {
	result = database.GetConnection().Debug().Updates(entity)
	err := result.Error
	if err != nil {
		log.Println("Error >>>", err.Error())
	}
	return result
}

func Delete(entity interface{}, id int) (result *gorm.DB) {
	result = database.GetConnection().Delete(entity, id)
	if result.Error != nil {
		log.Println("Error when delete data >>>", result.Error)
	}
	return
}
