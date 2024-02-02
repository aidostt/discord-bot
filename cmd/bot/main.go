package main

import (
	"discord-bot.aidostt.me/internal/bot"
	"discord-bot.aidostt.me/pkg/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	newBot, err := bot.NewBot(cfg)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	if err := newBot.Start(); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
	fmt.Println("bot is successfully started!")
	newBot.AwaitSignal()
}
