package utils

import (
	"fmt"
	"os"
)

func RadarrMessage(eventType string) string {
	var message string

	switch eventType {
	case "Grab":
		message = fmt.Sprintf("<b>🟠🔍 Radarr - Movie Grabbed</b>\n%s", os.Getenv("radarr_movie_title"))

	case "Download":
		if os.Getenv("radarr_isupgrade") == "True" {
			message = fmt.Sprintf("<b>🟠🆙 Radarr - Movie Upgraded</b>\n%s", os.Getenv("radarr_movie_title"))
		} else {
			message = fmt.Sprintf("<b>🟠☑️ Radarr - Movie Downloaded</b>\n%s", os.Getenv("radarr_movie_title"))
		}

	case "Rename":
		message = fmt.Sprintf("<b>🟠🛠️ Radarr - Movie Renamed</b>\n%s\n<code>%s</code>", os.Getenv("radarr_movie_title"), os.Getenv("radarr_movie_path"))

	case "MovieDelete":
		message = fmt.Sprintf("<b>🟠🗑️ Radarr - Movie Deleted</b>\n%s\n<code>%s</code>", os.Getenv("radarr_movie_title"), os.Getenv("radarr_movie_path"))

	case "HealthIssue":
		switch os.Getenv("radarr_health_issue_level") {
		case "Ok":
			message = fmt.Sprintf("<b>🟠💚 Radarr - Health OK</b>\n%s\n%s", os.Getenv("radarr_health_issue_level"), os.Getenv("radarr_health_issue_message"))
		case "Notice":
			message = fmt.Sprintf("<b>🟠💙 Radarr - Health Notice</b>\n%s\n%s", os.Getenv("radarr_health_issue_level"), os.Getenv("radarr_health_issue_message"))
		case "Warning":
			message = fmt.Sprintf("<b>🟠🧡 Radarr - Health Warning</b>\n%s\n%s", os.Getenv("radarr_health_issue_level"), os.Getenv("radarr_health_issue_message"))
		case "Error":
			message = fmt.Sprintf("<b>🟠💔 Radarr - Health Error</b>\n%s\n%s", os.Getenv("radarr_health_issue_level"), os.Getenv("radarr_health_issue_message"))
		}

	case "HealthRestored":
		message = fmt.Sprintf("<b>🟠💚 Radarr - Health Restored</b>\n%s\n%s", os.Getenv("radarr_health_issue_level"), os.Getenv("radarr_health_issue_message"))

	case "ApplicationUpdate":
		message = fmt.Sprintf("<b>🟠🔧 Radarr - Update</b>\n<code>%s -> %s</code>\n%s", os.Getenv("radarr_update_previousversion"), os.Getenv("radarr_update_newversion"), os.Getenv("radarr_update_message"))

	case "Test":
		message = "<b>🟠🔧 Radarr - Test</b>\nThis is a test notification."

	default:
		fmt.Printf("[WARN] Radarr - Unknown event type: %s\n", eventType)
		message = fmt.Sprintf("<b>🟠 Radarr - Error</b>\nEvent type: %s", eventType)
	}

	return message
}
