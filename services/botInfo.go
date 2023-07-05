package services

import (
	"alertbot/blockchain"
	"alertbot/notification"
	"alertbot/rpc"
	"alertbot/validator"

	"github.com/gofiber/fiber/v2"
)

func GetBotInfo(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"rpc":              rpc.RPC,
		"blockHeight":      blockchain.GetCurrentBlockHeight(),
		"validator":        validator.Validator,
		"currentSignature": blockchain.GetCurrentSignature(),
		"webhookUrl":       notification.GetWekhook(),
	})
}
