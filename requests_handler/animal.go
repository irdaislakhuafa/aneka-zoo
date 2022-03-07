package requests_handler

import (
	"aneka-zoo/entities"
	"aneka-zoo/helpers"
	"aneka-zoo/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewAnimal(reqAndRes *gin.Context) {
	var newAnimal entities.Animal
	response := make(map[string]interface{})
	var status int

	defer func() {
		if r := recover(); r != nil {
			response["data"] = nil
			response["message"] = r
			response["fields"] = helpers.GetFields(newAnimal)
			reqAndRes.JSON(http.StatusInternalServerError, response)
		}
	}()

	err := reqAndRes.ShouldBindJSON(&newAnimal)
	newAnimal.Name = strings.Trim(newAnimal.Name, " ")
	newAnimal.Class = strings.Trim(newAnimal.Class, " ")

	// if json format invalid with Animal struct
	if err != nil {
		status = http.StatusBadRequest
		response["data"] = nil
		response["message"] = "Please check your JSON format!"
		response["fields"] = helpers.GetFields(newAnimal) // get all fields "json" in struct
	} else {
		err = services.Create(&newAnimal).Error
		// if error when create new data in database
		if err != nil {
			response["data"] = nil

			// if error is duplicate entry
			if strings.Contains(strings.ToLower(err.Error()), "duplicate entry") {
				response["message"] = "data already exists!"
			} else {
				response["message"] = err.Error()
			}

			response["fields"] = helpers.GetFields(newAnimal)
		} else {
			// if JSON format is valid and no error when save data
			response["data"] = newAnimal
			response["message"] = "success saved data!"
			response["fields"] = helpers.GetFields(newAnimal)
		}
	}

	reqAndRes.JSON(
		status,
		response,
	)
}

func UpdateAnimal(reqAndRes *gin.Context) {
	animal := entities.Animal{}
	response := make(map[string]interface{})
	status := http.StatusOK
	response["fields"] = helpers.GetFields(animal)

	defer func() {
		if r := recover(); r != nil {
			status = http.StatusInternalServerError
			response["data"] = nil
			response["message"] = r
			reqAndRes.JSON(status, response)
		}
	}()

	// get id
	id, err := strconv.Atoi(reqAndRes.Param("id"))

	if err != nil {
		status = http.StatusBadRequest
		response["data"] = nil
		response["message"] = "ID must be int or number!"
	} else {
		err := reqAndRes.ShouldBindJSON(&animal)
		// FIXME :please finx me :"D
		if err != nil {
			status = http.StatusBadRequest
			response["data"] = nil
			response["message"] = "pleace check your JSON format!"
		} else {
			animal.ID = id
			if services.AnimalExistsById(id) {
				response["message"] = "success updated data!"
				services.Update(&animal)
			} else {
				services.Create(&animal)
				response["message"] = "data does not exists, success created new data!"
				// response["message"] = id
			}
			response["data"] = animal

		}

		/* if services.AnimalExistsById(id) {
			err = reqAndRes.ShouldBindJSON(&animal)
			if err != nil {
				status = http.StatusBadRequest
				response["data"] = nil
				response["message"] = "pleace check your JSON format!"
			} else {
				status = http.StatusOK
				animal.ID = id
				services.Update(&animal)
				response["data"] = animal
				response["message"] = "success updated data!"
			}
		} else {
			if err != nil {
				status = http.StatusBadRequest
				response["data"] = nil
				response["message"] = "pleace check your JSON format!"
			} else {
				animal.ID = id
				services.Update(&animal)
				response["message"] = "data does not exists, success created new data!"
				response["data"] = animal
			}
		} */
	}

	reqAndRes.JSON(status, response)
}

func DeleteAnimal(reqAndRes *gin.Context) {
	response := make(map[string]interface{})
	status := http.StatusOK
	var deletedAnimal entities.Animal

	// if internal server error
	defer func() {
		if r := recover(); r != nil {
			response["data"] = nil
			response["message"] = r
			status = http.StatusInternalServerError
			reqAndRes.JSON(status, response)
		}
	}()

	// convert id string to int
	id, err := strconv.Atoi(reqAndRes.Param("id"))

	switch {
	case err != nil: // if error is exists
		status = http.StatusBadRequest
		response["data"] = nil
		response["message"] = "ID must be int or number!"
	case err == nil:
		err = services.FindAnimalById(id, &deletedAnimal).Error
		if err != nil {
			status = http.StatusNotFound
			response["data"] = nil
			response["message"] = "data doesn't exists!"
		} else {
			status = http.StatusOK
			services.Delete(&entities.Animal{}, id)
			response["data"] = deletedAnimal
			response["message"] = "success deleted data!"
		}
	}

	reqAndRes.JSON(status, response)
}

func GetAnimalById(reqAndRes *gin.Context) {
	response := make(map[string]interface{})
	status := http.StatusOK

	id, err := strconv.Atoi(reqAndRes.Param("id"))

	if err != nil {
		status = http.StatusBadRequest
		response["data"] = nil
		response["message"] = "pleace check your JSON format!"
	} else {
		animal := entities.Animal{}
		if services.AnimalExistsById(id) {
			status = http.StatusOK
			services.FindAnimalById(id, &animal)
			response["data"] = animal
			response["message"] = "success get data!"
		} else {
			status = http.StatusNotFound
			response["data"] = nil
			response["message"] = "data not found!"
		}
	}

	reqAndRes.JSON(status, response)
}
func GetAllAnimals(reqAndRes *gin.Context) {
	response := make(map[string]interface{})
	status := http.StatusOK
	var animals []entities.Animal

	result := services.FindAll(&animals)

	if result.Error != nil {
		response["data"] = nil
		response["message"] = result.Error.Error()
	} else {
		response["data"] = animals
		response["message"] = "success find all data"
	}

	reqAndRes.JSON(status, response)
}
