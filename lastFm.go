package lastFm

import (
	"github.com/Cellularhacker/logger-go"
)

const (
	lastFmDefaultAPIEndpoint = "https://ws.audioscrobbler.com/2.0/"
)

var (
	initialized = false
)

var (
	apiEndpoint = ""
	apiKey      = ""
	apiSecret   = ""
)

// SetVariables - Mark: You should have to call this function while initializing step or before run main.
func SetVariables(lastFmApiKey, lastFmAPISecret string, lastFmAPIEndpoint ...string) {

	//	MARK: Check if it initialized
	if len(lastFmAPIEndpoint) <= 0 {
		logger.L.Warnf("'lastFmAPIEndpoint' is missing!")
		logger.L.Warnf("Using default API endpoint: %s", lastFmDefaultAPIEndpoint)
		apiEndpoint = lastFmDefaultAPIEndpoint
	} else {
		apiEndpoint = lastFmAPIEndpoint[0]
	}
	if lastFmApiKey == "" {
		logger.L.Fatalf("'lastFmApiKey' is missing!")
	}
	apiKey = lastFmApiKey
	if lastFmAPISecret == "" {
		logger.L.Fatalf("'lastFmAPISecret' is missing!")
	}
	apiSecret = lastFmAPISecret

	initialized = true
}

func IsInitialized() bool {
	return initialized
}
