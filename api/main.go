package main

import (
	"os"
	"project/CodeSancus-Backend/database"
	"project/CodeSancus-Backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.ConnectDatabase()

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	routes.DatasetRoutes(router)
	router.Run()

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:8000/upload"}
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	// config.AllowHeaders = []string{"Content-Type"}
	// config.AllowCredentials = true
	// router.Use(cors.New(config))

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8000/upload"},
	// 	AllowMethods:     []string{"POST", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length", "Content-Type"},
	// 	AllowCredentials: true,
	//   }))
}
