package main

import (
	"os"
	"project/CodeSancus-Backend/database"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.ConnectDatabase()
}
