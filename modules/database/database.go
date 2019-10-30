package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ConnectDB Function connects to the database
func ConnectDB() *gorm.DB {
	host := os.Getenv("dbHost")
	port := os.Getenv("dbPort")
	username := os.Getenv("dbUser")
	pass := os.Getenv("dbPass")
	dbname := os.Getenv("dbName")
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, pass, dbname)

	// Connect DB
	db, err := gorm.Open("postgres", dbInfo) // assign postgres driver
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected established successfully to %s through port %s\n", host, port)
	return db
}
