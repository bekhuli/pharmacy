package main

import (
	"log"

	"github.com/bekhuli/pharmacy/pkg/db"
)

func main() {
	db.Connect()

	defer func() {
		log.Println("Closing database connection")
		if err := db.DB.Close(); err != nil {
			log.Println("Failed to close database connection:", err)
		} else {
			log.Println("Database disconnected successfully")
		}
	}()
}
