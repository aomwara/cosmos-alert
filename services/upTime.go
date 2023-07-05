package services

import (
	"alertbot/validator"

	"github.com/gofiber/fiber/v2"
)

func GetUptime(c *fiber.Ctx) error {
	uptime := validator.GetUptime()
	return c.JSON(fiber.Map{
		"uptime": uptime,
	})
}
