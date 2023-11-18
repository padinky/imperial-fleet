package helper

import "github.com/gofiber/fiber/v2"

func ResponseNotFound(c *fiber.Ctx, message string) error {
	return c.Status(404).JSON(fiber.Map{
		"code":    404,
		"message": message,
	})
}

func ResponseBadRequest(c *fiber.Ctx, message string) error {
	return c.Status(400).JSON(fiber.Map{
		"code":    400,
		"message": message,
	})
}

func ResponseUnauthorized(c *fiber.Ctx, message string) error {
	return c.Status(401).JSON(fiber.Map{
		"code":    401,
		"message": message,
	})
}

func ResponseError(c *fiber.Ctx, message string) error {
	return c.Status(500).JSON(fiber.Map{
		"code":    500,
		"message": message,
	})
}

func ResponseSuccessWithCodeAndData(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

func ResponseSuccessOnly(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
	})
}

func ResponseDataOnly(c *fiber.Ctx, data interface{}) error {
	return c.JSON(data)
}
