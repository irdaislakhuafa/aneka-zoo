package requests_handler

import (
	"aneka-zoo/entities"
	"aneka-zoo/helpers"
	"aneka-zoo/services"
	"net/http"
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
