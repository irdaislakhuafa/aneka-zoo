package services

import (
	"aneka-zoo/database"
	"aneka-zoo/entities"
	"log"

	"gorm.io/gorm"
)

func FindAnimalById(id int, entity *entities.Animal) *gorm.DB {
	result := database.GetConnection().First(entity, id)
	err := result.Error
	if err != nil {
		log.Println("Error >>>", err.Error)
	}
	return result
}
func AnimalExistsById(id int) bool {
	effected, _ := FindAnimalById(id, &entities.Animal{}).Rows()
	return effected != nil
}
