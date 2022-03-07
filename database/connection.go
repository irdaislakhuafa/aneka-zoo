package database

import "gorm.io/gorm"

var connection *gorm.DB

func SetConnection(con *gorm.DB) {
	connection = con
}
func GetConnection() *gorm.DB {
	return connection
}
