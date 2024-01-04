package pkg

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"
	"weatherbot/internal"

	tgapi "git.foxminded.ua/foxstudent106270/telegramapi.git"
	"github.com/phuslu/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHTTPClient struct {
	mock.Mock
}

func (c *MockHTTPClient) PostForm(url string, data url.Values) (*http.Response, error) {
	args := c.Called(url, data)
	return args.Get(0).(*http.Response), args.Error(1)
}

func requestBody(t *testing.T, country string) *tgapi.WebHookReqBody {
	reqBody := &tgapi.WebHookReqBody{
		Message: tgapi.Message{
			Text: country,
			Chat: tgapi.Chat{
				ID: 358383178,
			},
		},
	}
	return reqBody
}

func TestHandlerTelegram(t *testing.T) {
	log.DefaultLogger = log.Logger{
		Level:      log.DebugLevel,
		Caller:     1,
		TimeFormat: time.RFC850,
		Writer:     &log.ConsoleWriter{},
	}

	holidayAPI := GetHolidayAPI()

	tests := []struct {
		name     string
		country  string
		have     url.Values
		expected url.Values
	}{
		{
			name:    "invalid country",
			country: "AAA",
			have: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"No holiday today in AAA"},
			},
			expected: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"No holiday today in AAA"},
			},
		},

		{
			name:    "US",
			country: string(internal.CountryFlags["US"]),
			have: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"Today is Hotdog day in US"},
			},
			expected: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"Today is Hotdog day in US"},
			},
		},

		{
			name:    "IT",
			country: string(internal.CountryFlags["IT"]),
			have: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"Today is Pizza day in IT"},
			},
			expected: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"Today is Pizza day in IT"},
			},
		},
		{
			name:    "JP",
			country: string(internal.CountryFlags["JP"]),
			have: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"Today is Sushi day in JP"},
			},
			expected: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"Today is Sushi day in JP"},
			},
		},
		{
			name:    "DE",
			country: string(internal.CountryFlags["DE"]),
			have: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"Today is Beer day in DE"},
			},
			expected: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"Today is Beer day in DE"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			holidayAPI.HolidayRequest(requestBody(t, tt.country), 358383178)
			assert.Equal(t, tt.have, tt.expected)
		})
	}
}
func Test_getEvent(t *testing.T) {
	type args struct {
		events  []Event
		chatID  int
		country string
	}

	tests := []struct {
		name string
		args args
		want url.Values
	}{
		{
			name: "0 holidays",
			args: args{
				events:  []Event{},
				chatID:  358383178,
				country: "JP",
			},
			want: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {"No holiday today in JP"},
			},
		},
		{
			name: "1 holiday",
			args: args{
				events:  []Event{{Name: "New Year"}},
				chatID:  358383178,
				country: "US",
			},
			want: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {fmt.Sprintf("Today is %s in %s", "New Year", "US")},
			},
		},

		{
			name: "5 holidays",
			args: args{
				events:  []Event{{Name: "New Year"}, {Name: "Christmas"}, {Name: "Whatever holiday"}, {Name: "Man Utd won"}, {Name: "Pizza day"}},
				chatID:  358383178,
				country: "JP",
			},
			want: url.Values{
				"chat_id": {strconv.Itoa(358383178)},
				"text":    {fmt.Sprintf("Today is %s, %s, %s, %s, %s in %s", "New Year", "Christmas", "Whatever holiday", "Man Utd won", "Pizza day", "JP")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event, _ := getEvent(tt.args.events, tt.args.chatID, tt.args.country)
			assert.Equalf(t, tt.want, event, "getEvent(%v, %v, %v)", tt.args.events, tt.args.chatID, tt.args.country)
		})
	}
}
