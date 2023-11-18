package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/padinky/imperial-fleet/database"
	"github.com/padinky/imperial-fleet/handlers"
	"github.com/padinky/imperial-fleet/middleware"
	"github.com/padinky/imperial-fleet/repository"
)

func Initalize(router *fiber.App) {

	db := database.DB
	router.Use(middleware.ApplySecurity)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	spaceshipRepo := repository.NewSpaceshipRepository()
	spaceshipHandler := handlers.NewSpaceshipHandler(spaceshipRepo)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	spaceshipRoute := v1.Group("/spaceship")
	spaceshipRoute.Post("/", middleware.Authenticated, spaceshipHandler.Create)
	spaceshipRoute.Get("/:id", spaceshipHandler.GetByID)
	spaceshipRoute.Get("/", spaceshipHandler.GetAll)
	spaceshipRoute.Put("/:id", middleware.Authenticated, spaceshipHandler.Update)
	spaceshipRoute.Delete("/:id", middleware.Authenticated, spaceshipHandler.Delete)

	authHandler := handlers.NewAuthHandler(db)

	users := v1.Group("/users")
	users.Post("/", authHandler.CreateUser)
	users.Post("/login", authHandler.Login)
	users.Delete("/logout", middleware.Authenticated, authHandler.Logout)
	users.Delete("/", middleware.Authenticated, authHandler.DeleteUser)
	users.Put("/", middleware.Authenticated, authHandler.ChangePassword)
	users.Get("/me", middleware.Authenticated, authHandler.GetUserInfo)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
