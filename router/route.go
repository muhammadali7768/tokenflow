package router

import (
	"example.com/go-fiber/controller"

	"example.com/go-fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	auth := api.Group("/auth", logger.New())
	auth.Post("/register", controller.RegisterUser)
	auth.Post("/login", controller.LoginUser)
}

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", middleware.AuthMiddleware, logger.New())

	// routes
	api.Get("/", controller.GetAllProducts)
	api.Get("/:id", controller.GetSingleProduct)
	api.Post("/", controller.CreateProduct)
	api.Delete("/:id", controller.DeleteProduct)

	api.Get("/balance/:address", controller.GetBalance)
	api.Post("/send-balance", controller.SendBalanceToAddress)
}
