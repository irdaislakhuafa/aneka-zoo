package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

func SetConnection(con *gorm.DB) {
	connection = con
}
func GetConnection() *gorm.DB {
	return connection
}
func OpenConnection(username, password, dbHost, dbPort, dbName string) {
	dbUrl := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal("Error create connection to database :", err.Error())
	} else {
		SetConnection(db)
		log.Println("Succes create connection to database :", dbName)
	}
}
