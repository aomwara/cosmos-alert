package services

import (
	"alertbot/blockchain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCurentBlock(c *fiber.Ctx) error {
	currentBlock := blockchain.GetCurrentBlockHeight()
	return c.SendString(strconv.Itoa(currentBlock))
}
