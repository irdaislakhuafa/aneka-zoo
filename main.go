package main

import (
	"aneka-zoo/database"
	"aneka-zoo/entities"
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dbUrl := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Error create connection to database :", err.Error())
	} else {
		database.SetConnection(db)
		log.Println("Succes create connection to database :", dbName)
	}

	// auto migrate
	err = nil
	err = database.GetConnection().AutoMigrate(
		&entities.Animal{},
	)
	if err != nil {
		log.Fatal("Error when migrate to database :", err.Error())
	} else {
		database.SetConnection(db)
		log.Println("Succes create migration model to database :", dbName)
	}

	// create router gin
	router := gin.Default()

	// run router
	appPort = ":" + appPort
	router.Run(appPort)
}
