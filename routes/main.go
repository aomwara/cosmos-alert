package routes

import (
	"alertbot/services"

	"github.com/gofiber/fiber/v2"
)

func Handle(app *fiber.App) {
	app.Get("/", services.GetHome)
	app.Get("/currentBlock", services.GetCurentBlock)
	app.Get("/bot-info", services.GetBotInfo)
	app.Get("/uptime", services.GetUptime)

}
