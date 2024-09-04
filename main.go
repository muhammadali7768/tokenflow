package main

import (
	"log"

	"example.com/go-fiber/config"
	"example.com/go-fiber/database"
	"example.com/go-fiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/lib/pq"
)

// entry point to our program
func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatalf("Database connection Error: %v ", err)
	}

	// call the New() method - used to instantiate a new Fiber App
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// Initialize the session store
	config.InitSessionStore()

	// Use the session middleware
	app.Use(config.SessionMiddleware)

	router.RegisterRoutes(app)
	router.SetupRoutes(app)

	// listen on port 3000
	app.Listen(":3000")

}
