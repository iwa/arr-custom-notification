package main

import (
	"fmt"
	"os"

	"github.com/iwa/arr-custom-notification/internal/utils"
)

func main() {
	botToken, chatId := utils.ImportEnv()

	sonarrEventType := os.Getenv("sonarr_eventtype")
	radarrEventType := os.Getenv("radarr_eventtype")

	if sonarrEventType == "" && radarrEventType == "" {
		panic("[FATAL] No event type found.")
	}

	var message string

	if sonarrEventType != "" {
		fmt.Println("Sonarr - Event type:", sonarrEventType)
		message = utils.SonarrMessage(sonarrEventType)
	} else if radarrEventType != "" {
		fmt.Println("Radarr - Event type:", radarrEventType)
		message = utils.RadarrMessage(radarrEventType)
	}

	err := utils.SendTelegramMessage(chatId, botToken, message)
	if err != nil {
		fmt.Printf("Error sending Telegram message: %s\n", err)
		panic("[FATAL] Failed to send Telegram message.")
	}
}
