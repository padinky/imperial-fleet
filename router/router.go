package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/padinky/imperial-fleet/handlers"
	"github.com/padinky/imperial-fleet/repository"
)

func Initalize(router *fiber.App) {

	// db := database.DB
	// router.Use(middleware.ApplySecurity)

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	spaceshipRepo := repository.NewSpaceshipRepository()
	spaceshipHandler := handlers.NewSpaceshipHandler(spaceshipRepo)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	spaceshipRoute := v1.Group("/spaceship")
	spaceshipRoute.Post("/", spaceshipHandler.Create)
	spaceshipRoute.Get("/:id", spaceshipHandler.GetByID)
	spaceshipRoute.Get("/", spaceshipHandler.GetAll)
	spaceshipRoute.Put("/:id", spaceshipHandler.Update)
	spaceshipRoute.Delete("/:id", spaceshipHandler.Delete)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
