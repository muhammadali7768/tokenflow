package router

import (
	"fmt"

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
	//api.Get("/:id", controller.GetSingleProduct)
	//api.Post("/", controller.CreateProduct)
	//api.Delete("/:id", controller.DeleteProduct)

	api.Get("/balance/:address", controller.GetBalance)
	api.Get("/engc-balance/:address", controller.GetEngageCBalance)
	api.Post("/send-balance", controller.SendBalanceToAddress)
	api.Post("/deploy-engage-coin-contract", controller.DeployEngageCoin)
	api.Get("/get-recent-transactions", controller.GetRecentTransactions)
	api.Post("/send-reward", controller.TransferReward)
	api.Get("/get-engc-token-distribution", controller.GetEngcTokenDistribution)
	api.Post("/deploy-engc-stacking-contract", controller.DeployENGCStaking)
	api.Post("/stake-engc", controller.DepositAmount)
	api.Get("/stake/:address", controller.GetStakedAmount)
	api.Get("/total-staked", controller.GetTotalStacked)

	api.Get("/test-user-id", func(c *fiber.Ctx) error {
		userId, err := controller.GetUserId(c)
		fmt.Print("UserID", userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"success": true,
			"userId":  userId,
		})
	})
}
