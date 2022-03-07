package main

import (
	"aneka-zoo/database"
	"aneka-zoo/entities"
	"aneka-zoo/requests_handler"
	"aneka-zoo/services"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var username, password, dbName, dbPort, dbHost, appPort string
	// clear scree
	fmt.Println("\033\143")

	// get value option from command line
	flag.StringVar(&username, "username", "root", "Put your database username here!")        // username
	flag.StringVar(&password, "password", "", "Put your database password here!")            // password
	flag.StringVar(&dbName, "db-name", "aneka_zoo", "Put your database name here!")          // dbName
	flag.StringVar(&dbPort, "db-port", "3306", "Put your database port here!")               // dbPort
	flag.StringVar(&dbHost, "db-host", "127.0.0.1", "Put your database host here!")          // dbPort
	flag.StringVar(&appPort, "app-port", "8080", "Put your database application port here!") // appPort
	flag.Parse()

	// open database connection
	database.OpenConnection(username, password, dbHost, dbPort, dbName)

	// auto migrate
	services.AutoMigration(&entities.Animal{})

	// create router gin
	router := gin.Default()
	api := router.Group("/api")

	// apis router
	api.POST("/animal/new", requests_handler.NewAnimal)
	api.PUT("/animal/update/:id", requests_handler.UpdateAnimal)

	// run router
	appPort = ":" + appPort
	router.Run(appPort)
}
