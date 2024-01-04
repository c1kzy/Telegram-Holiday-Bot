package pkg

// HolidayConfig struct for Holiday API
type HolidayConfig struct {
	API string `env:"API"`
}

// WebHookReqBody struct for telegram body response
type WebHookReqBody struct {
	Message  Message  `json:"message"`
	Location Location `json:"location"`
}

// Message struct with text and chat
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	Date int64  `json:"date"`
}

// Chat struct with chat ID
type Chat struct {
	ID int `json:"id"`
}

// Location struct for telegram body
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
