package internal

// StartText for telegram /start command
const StartText = `Hello! Please choose one of the following countries to get today's holiday in that country
`

// Flag type creates new type Flag
type Flag string

// Flag constants
const (
	USFlag Flag = "\xF0\x9F\x87\xBA\xF0\x9F\x87\xB8"
	ITFlag Flag = "\xF0\x9F\x87\xAE\xF0\x9F\x87\xB9"
	DEFlag Flag = "\xF0\x9F\x87\xA9\xF0\x9F\x87\xAA"
	JPFlag Flag = "\xF0\x9F\x87\xAF\xF0\x9F\x87\xB5"
)

// Flags map for emoji and flags for telegram
var Flags = map[Flag]string{
	"\xF0\x9F\x87\xBA\xF0\x9F\x87\xB8": "US",
	"\xF0\x9F\x87\xA9\xF0\x9F\x87\xAA": "DE",
	"\xF0\x9F\x87\xAE\xF0\x9F\x87\xB9": "IT",
	"\xF0\x9F\x87\xAF\xF0\x9F\x87\xB5": "JP",
}

// CountryFlags map for Telegram buttons
var CountryFlags = map[string]Flag{
	"US": USFlag,
	"DE": DEFlag,
	"IT": ITFlag,
	"JP": JPFlag,
}
