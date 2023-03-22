package variable

import "errors"

var (
	ErrTwitchInvalid = errors.New("Twitch: Invalid response from Twitch API")
	JSONParseError   = errors.New("JSON: Error parsing JSON")
	FailToInitialize = errors.New("Twitch: Failed to initialize Twitch client")
)
