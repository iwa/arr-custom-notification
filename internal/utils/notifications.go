package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type TelegramMessage struct {
	ChatID              string `json:"chat_id"`
	Text                string `json:"text"`
	ParseMode           string `json:"parse_mode"`
	DisableNotification bool   `json:"disable_notification"`
	ProtectContent      bool   `json:"protect_content"`
}

func SendTelegramMessage(chatId, botToken, message string) error {
	payload := TelegramMessage{
		ChatID:              chatId,
		Text:                message,
		ParseMode:           "HTML",
		DisableNotification: true,
		ProtectContent:      false,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response from Telegram API:", string(responseBody))

	return nil
}

// Function to format episodes
func formatEpisodes(seasonNumber, episodeNumbers, episodeTitles string) (string, error) {
	numbers := strings.Split(episodeNumbers, ",")
	titles := strings.Split(episodeTitles, "|")

	if len(numbers) != len(titles) {
		return "", fmt.Errorf("mismatch between episode numbers and titles")
	}

	var formattedEpisodes strings.Builder
	for i := range numbers {
		num, err := strconv.Atoi(numbers[i])
		if err != nil {
			return "", fmt.Errorf("failed to convert episode number to int: %w", err)
		}

		if i != 0 {
			formattedEpisodes.WriteString("\n")
		}

		formattedEpisodes.WriteString(fmt.Sprintf("<code>%sx%02d</code> - %s", seasonNumber, num, titles[i]))
	}

	return formattedEpisodes.String(), nil
}
