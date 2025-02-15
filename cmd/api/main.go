package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using default env variables")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Spy Cat Agency API is running..."})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	fmt.Println("Starting Spy Cat Agency API on port " + port + "...")
	r.Run(":" + port)
}
