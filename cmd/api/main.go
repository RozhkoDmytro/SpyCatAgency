package main

import (
	"log"
	"os"

	"github.com/RozhkoDmytro/SpyCatAgency/internal/config"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/middleware"
	"github.com/RozhkoDmytro/SpyCatAgency/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	logFile, err := middleware.InitMiddlewareLogger()
	if err != nil {
		log.Fatal("❌ Failed to initialize middleware logger:", err)
	}
	defer logFile.Close()

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Warning: No .env file found, using default env variables")
	}

	db := config.InitDB()

	r := router.InitRouter(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Spy Cat Agency API is running on port " + port + "...")
	r.Run(":" + port)
}
