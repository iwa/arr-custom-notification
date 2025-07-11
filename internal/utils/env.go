package utils

import "os"

func ImportEnv() (string, string) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		panic("[ERROR] No domains provided. Please set the DOMAINS environment variable as comma-separated values.")
	}

	chatId := os.Getenv("TELEGRAM_CHAT_ID")
	if chatId == "" {
		panic("[ERROR] No chat ID provided. Please set the TELEGRAM_CHAT_ID environment variable.")
	}

	return botToken, chatId
}
