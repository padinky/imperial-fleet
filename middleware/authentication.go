package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/padinky/imperial-fleet/database"
	"github.com/padinky/imperial-fleet/handlers"
	"github.com/padinky/imperial-fleet/helper"
	"github.com/padinky/imperial-fleet/model"
)

func Authenticated(c *fiber.Ctx) error {
	cookie := new(model.Session)
	if err := c.CookieParser(cookie); err != nil {
		return helper.ResponseUnauthorized(c, "unauthorized request")
	}

	db := database.DB
	authHandler := handlers.NewAuthHandler(db)
	user, err := authHandler.GetUser(cookie.SessionID)
	if err != nil {
		return helper.ResponseUnauthorized(c, "unauthorized request")
	}
	c.Locals("user", user)
	return c.Next()
}
