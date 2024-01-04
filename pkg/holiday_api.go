package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"weatherbot/internal"

	api "github.com/c1kzy/Telegram-API"

	"github.com/phuslu/log"
	"github.com/zsefvlol/timezonemapper"
)

// HolidayAPI struct for holiday config
type HolidayAPI struct {
	API string
}

var (
	lock        = sync.Mutex{}
	sHolidayAPI *HolidayAPI
	events      []Event
)

// Event struct for holiday names
type Event struct {
	Name string `json:"name"`
}

// GetHolidayAPI is getting single instance for API
func GetHolidayAPI() *HolidayAPI {
	if sHolidayAPI == nil {
		lock.Lock()
		defer lock.Unlock()
		if sHolidayAPI == nil {
			sHolidayAPI = &HolidayAPI{
				API: os.Getenv("API"),
			}
			log.Info().Msg("Holiday API created")
		}
	}

	return sHolidayAPI
}

// HolidayRequest function for API request getting today's holiday
func (h *HolidayAPI) HolidayRequest(body *api.WebHookReqBody, chatID int) (url.Values, error) {
	country, found := internal.Flags[internal.Flag(body.Message.Text)]
	if !found {
		return url.Values{}, fmt.Errorf("country was not found in available country list")
	}

	timeZone := timezonemapper.LatLngToTimezoneString(body.Message.Location.Latitude, body.Message.Location.Longitude)

	location, locError := time.LoadLocation(timeZone)
	if locError != nil {
		log.Error().Err(fmt.Errorf("unable to get time zone for holiday API: %w", locError))
	}

	currentTime := time.Now().In(location)

	resp, err := http.Get(fmt.Sprintf("https://holidays.abstractapi.com/v1/?api_key=%v&country=%s&year=%v&month=%v&day=%v", h.API, country, currentTime.Year(), int(currentTime.Month()), currentTime.Day()))
	if err != nil {
		return url.Values{}, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return url.Values{}, err
	}

	marshalErr := json.Unmarshal(respBody, &events)
	if marshalErr != nil {
		log.Error().Err(fmt.Errorf("error unmarshalling JSON for holiday API request Error:%w", err))
	}

	return getEvent(events, chatID, country)
}
func getEvent(events []Event, chatID int, country string) (url.Values, error) {
	if len(events) == 0 {
		return url.Values{
			"chat_id": {strconv.Itoa(chatID)},
			"text":    {fmt.Sprintf("There are no holidays today in %s", country)},
		}, nil
	}

	holidayBuilder := strings.Builder{}

	for i, event := range events {
		if i > 0 {
			holidayBuilder.WriteString(", ")
		}
		holidayBuilder.WriteString(event.Name)
	}

	return url.Values{
		"chat_id": {strconv.Itoa(chatID)},
		"text":    {fmt.Sprintf("There is %s today in %s", holidayBuilder.String(), country)},
	}, nil

}
