package internal

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	api "github.com/c1kzy/Telegram-API"
)

// KeyboardButton struct for button text
type KeyboardButton struct {
	Text     string `json:"text"`
	Location bool   `json:"request_location"`
}

// ReplyKeyboardMarkup struct for keyboard layout
type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"`
}

// StartResponse function for /start command
func StartResponse(body *api.WebHookReqBody, chatID int) (url.Values, error) {
	replyMarkup := ReplyKeyboardMarkup{
		Keyboard: [][]KeyboardButton{
			{KeyboardButton{Text: string(CountryFlags["US"])}, KeyboardButton{Text: string(CountryFlags["IT"])}},
			{KeyboardButton{Text: string(CountryFlags["DE"])}, KeyboardButton{Text: string(CountryFlags["JP"])}},
			{KeyboardButton{Text: "Share my location", Location: true}},
		},
	}
	jsonData, jsonErr := json.Marshal(replyMarkup)
	if jsonErr != nil {
		return url.Values{}, fmt.Errorf("error marshaling JSON: %w", jsonErr)
	}

	return url.Values{
		"chat_id":      {strconv.Itoa(chatID)},
		"text":         {StartText},
		"reply_markup": {string(jsonData)},
	}, nil
}
