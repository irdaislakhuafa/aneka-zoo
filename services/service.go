package services

import (
	"aneka-zoo/database"
	"log"
)

func Create(entity interface{}) {
	err := database.GetConnection().Create(entity).Error
	if err != nil {
		log.Println("Error created data >>>", err)
	}
}
