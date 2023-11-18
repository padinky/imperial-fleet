package main

import (
	"log"
	"os"

	"github.com/padinky/imperial-fleet/database"
	"github.com/padinky/imperial-fleet/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Custom File Writer
	file, err := os.OpenFile("./log/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output: file,
		Format: "${time} \t| ${status} \t| ${latency} \t| ${method} \t| ${path} \t| ${body} \t| ${resBody}\n",
	}))

	database.ConnectDB()

	router.Initalize(app)
	log.Fatal(app.Listen(":" + getenv("SERVICE_PORT", "3000")))
}
