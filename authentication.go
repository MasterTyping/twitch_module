package twitchmodule

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/MasterTyping/twitch_module/variable"
	"github.com/gobuffalo/logger"
)

type TwitchClient struct {
	client_id     string
	client_secret string
	grant_type    string
	logger        *logger.Logger
}

type TwitchUser struct {
}

func (t *TwitchClient) init(client_id string, client_secret string, grant_type string) {
	t.client_id = client_id
	t.client_secret = client_secret
	t.grant_type = grant_type

	var logger *logger.Logger
	t.logger = logger
}

var lock = &sync.Mutex{}

var Instance *TwitchClient

func initialize() (*TwitchClient, error) {
	if Instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if Instance == nil {
			Instance = &TwitchClient{}
			Instance.init("client_id", "client_secret", "grant_type")
		}
	}
	return Instance, nil
}

func (t *TwitchClient) GetTwitchAuthenticationToken(payload string) (*TwitchUser, error) {
	method := "POST"
	url := "https://id.twitch.tv/oauth2/token"

	err := json.Unmarshal([]byte(payload), &t)

	if err != nil {
		return nil, variable.JSONParseError
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))

	req.Header.Set("Content-Type", "application/json")

	return nil, nil
}

func WebHandler(w http.ResponseWriter, r *http.Request) {
	var t TwitchClient
	t.GetTwitchAuthenticationToken("payload")
}
