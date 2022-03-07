package services

import (
	"aneka-zoo/database"
	"log"
)

func AutoMigration(data interface{}) {
	err := database.GetConnection().AutoMigrate(
		data,
	)
	if err != nil {
		log.Fatal("Error when migrate to database :", err.Error())
	} else {
		log.Println("Succes create migration model to database :", database.GetConnection().Name())
	}
}
