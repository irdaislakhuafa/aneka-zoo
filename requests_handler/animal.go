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

	if err != nil {
		status = http.StatusBadRequest
		response["data"] = nil
		response["message"] = "Please check your JSON format!"
		response["fields"] = helpers.GetFields(newAnimal)
	} else {
		err = services.Create(&newAnimal).Error
		if err != nil {
			response["data"] = nil

			if strings.Contains(strings.ToLower(err.Error()), "duplicate entry") {
				response["message"] = "data already exists!"
			} else {
				response["message"] = err.Error()
			}

			response["fields"] = helpers.GetFields(newAnimal)
		} else {
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
