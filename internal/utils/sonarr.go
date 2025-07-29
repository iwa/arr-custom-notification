package utils

import (
	"fmt"
	"os"
)

func SonarrMessage(eventType string) string {
	var message string

	switch eventType {
	case "Grab":
		formattedEpisodes, err := formatEpisodes(os.Getenv("sonarr_release_seasonnumber"), os.Getenv("sonarr_release_episodenumbers"), os.Getenv("sonarr_release_episodetitles"))
		if err != nil {
			fmt.Println(eventType, err)
			formattedEpisodes = err.Error()
		}
		message = fmt.Sprintf("<b>🔵🔍 Sonarr - Episode Grabbed</b>\n%s\n%s", os.Getenv("sonarr_series_title"), formattedEpisodes)

	case "Download":
		formattedEpisodes, err := formatEpisodes(os.Getenv("sonarr_episodefile_seasonnumber"), os.Getenv("sonarr_episodefile_episodenumbers"), os.Getenv("sonarr_episodefile_episodetitles"))
		if err != nil {
			fmt.Println(eventType, err)
			formattedEpisodes = err.Error()
		}

		if os.Getenv("sonarr_isupgrade") == "True" {
			message = fmt.Sprintf("<b>🔵🆙 Sonarr - Episode Upgraded</b>\n%s\n%s", os.Getenv("sonarr_series_title"), formattedEpisodes)
		} else {
			message = fmt.Sprintf("<b>🔵☑️ Sonarr - Episode Downloaded</b>\n%s\n%s", os.Getenv("sonarr_series_title"), formattedEpisodes)
		}

	case "Rename":
		message = fmt.Sprintf("<b>🔵🛠️ Sonarr - Episode Renamed</b>\n%s\n<code>%s</code>", os.Getenv("sonarr_series_title"), os.Getenv("sonarr_series_path"))

	case "EpisodeFileDelete":
		formattedEpisodes, err := formatEpisodes(os.Getenv("sonarr_episodefile_seasonnumber"), os.Getenv("sonarr_episodefile_episodenumbers"), os.Getenv("sonarr_episodefile_episodetitles"))
		if err != nil {
			fmt.Println(eventType, err)
			formattedEpisodes = err.Error()
		}

		message = fmt.Sprintf("<b>🔵🗑️ Sonarr - Episode Deleted</b>\n%s\n%s", os.Getenv("sonarr_series_title"), formattedEpisodes)

	case "SeriesDelete":
		message = fmt.Sprintf("<b>🔵🗑️ Sonarr - Serie Deleted</b>\n%s\n<code>%s</code>", os.Getenv("sonarr_series_title"), os.Getenv("sonarr_series_path"))

	case "HealthIssue":
		switch os.Getenv("sonarr_health_issue_level") {
		case "Ok":
			message = fmt.Sprintf("<b>🔵💚 Sonarr - Health OK</b>\n%s\n%s", os.Getenv("sonarr_health_issue_type"), os.Getenv("sonarr_health_issue_message"))
		case "Notice":
			message = fmt.Sprintf("<b>🔵💙 Sonarr - Health Notice</b>\n%s\n%s", os.Getenv("sonarr_health_issue_type"), os.Getenv("sonarr_health_issue_message"))
		case "Warning":
			message = fmt.Sprintf("<b>🔵🧡 Sonarr - Health Warning</b>\n%s\n%s", os.Getenv("sonarr_health_issue_type"), os.Getenv("sonarr_health_issue_message"))
		case "Error":
			message = fmt.Sprintf("<b>🔵💔 Sonarr - Health Error</b>\n%s\n%s", os.Getenv("sonarr_health_issue_type"), os.Getenv("sonarr_health_issue_message"))
		}

	case "HealthRestored":
		message = fmt.Sprintf("<b>🔵💚 Sonarr - Health Restored</b>\n%s\n%s", os.Getenv("sonarr_health_issue_type"), os.Getenv("sonarr_health_issue_message"))

	case "ApplicationUpdate":
		message = fmt.Sprintf("<b>🔵🔧 Sonarr - Update</b>\n<code>%s -> %s</code>\n%s", os.Getenv("sonarr_update_previousversion"), os.Getenv("sonarr_update_newversion"), os.Getenv("sonarr_update_message"))

	case "Test":
		message = "<b>🔵🔧 Sonarr - Test</b>\nThis is a test notification."

	default:
		fmt.Printf("[WARN] Sonarr - Unknown event type: %s\n", eventType)
		message = fmt.Sprintf("<b>🔵 Sonarr - Error</b>\nEvent type: %s", eventType)
	}

	return message
}
