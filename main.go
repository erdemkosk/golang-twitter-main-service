package main

import (
	"log"
	"os"

	router "github.com/erdemkosk/golang-twitter-main-service/internal/routes"
	"github.com/erdemkosk/golang-twitter-main-service/pkg/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {

	godotenv.Load()
	database.ConnectDB()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // comma string format e.g. "localhost"
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Initalize(app)
	log.Fatal(app.Listen(":" + getenv("PORT", "3000")))
}
