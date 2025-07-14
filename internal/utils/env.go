package utils

import "os"

func ImportEnv() (string, string) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		panic("[ERROR] No bot token provided. Please set the TELEGRAM_BOT_TOKEN environment variable.")
	}

	chatId := os.Getenv("TELEGRAM_CHAT_ID")
	if chatId == "" {
		panic("[ERROR] No chat ID provided. Please set the TELEGRAM_CHAT_ID environment variable.")
	}

	return botToken, chatId
}
