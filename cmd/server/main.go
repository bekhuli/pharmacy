package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bekhuli/pharmacy/internal/common"
	"github.com/bekhuli/pharmacy/internal/routes"
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

	router := routes.InitRouter(db.DB)

	addr := fmt.Sprintf("%s:%s", common.ServerEnv.Host, common.ServerEnv.Port)
	log.Println("Server is running on port:", common.ServerEnv.Port)
	log.Fatal(http.ListenAndServe(addr, router))
}
