package main

import (
	"alertbot/blockchain"
	"alertbot/notification"
	"alertbot/routes"
	"alertbot/rpc"
	"alertbot/validator"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	args := os.Args
	args = args[1:]

	if len(args) != 3 {
		fmt.Println("Usage: alertbot <validator address> <rpc address> <discord_webhook>")
		return
	}

	//initial blockchain
	validator.SetValidator(args[0])
	rpc.SetRPC(args[1])
	notification.SetWebhook(args[2])
	blockchain.InitBlock()

	// Alert
	notification.Alert("Bot is running... ``` Validator: " + validator.Validator + "\n RPC: " + rpc.RPC + "\n Webhook: " + notification.GetWekhook() + "```\n")

	// Start Fiber API in a goroutine
	go func() {
		app := fiber.New()
		routes.Handle(app)
		app.Listen(":8000")
	}()

	var blockTime int = 2
	var stack int = 0

	var wg sync.WaitGroup

	for {
		wg.Add(1)
		go func(blockHeight int) {
			defer wg.Done()
			for {
				signature, err := rpc.GetSignature(blockHeight)
				if err != nil {
					stack++
					if stack == 0 {
						fmt.Println("Block: ", blockchain.GetCurrentBlockHeight(), " -> Retrying...")
						continue
					}

				} else {
					stack = 0
					fmt.Printf("Block Height: %d, Signature: %s\n", blockHeight, signature)
					if signature == "" {
						validator.IncreaseMissBlock()
						validator.SetUptime(false)
						fmt.Println("Block: ", blockchain.GetCurrentBlockHeight(), " -> Miss Block...")
					} else {
						validator.SetUptime(true)
						validator.DecreaseMissBlock()
					}
				}

				blockchain.SetCurrentSignature(signature)
				break
			}
		}(blockchain.GetCurrentBlockHeight())

		time.Sleep(time.Duration(blockTime) * time.Second)

		if stack == 0 {
			blockchain.IncreaseCurrentBlock()
		}

	}

}
