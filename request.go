package lastFm

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	KeyMethod = "method"
	KeyApiKey = "api_key"
)

func MakeRequest(method string, p url.Values, body io.Reader) ([]byte, error) {
	if !IsInitialized() {
		return nil, ErrNotInitialized
	}

	uri := apiEndpoint
	if len(p) <= 0 {
		p = url.Values{}
	}
	if len(p.Get(KeyMethod)) <= 0 {
		return nil, ErrNoMethodSpecified
	}
	p.Set(KeyApiKey, apiKey)

	uri = fmt.Sprintf("%s?%s", uri, p.Encode())

	req, _ := http.NewRequest(strings.ToUpper(method), uri, body)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do(req): %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 399 {
		respBody, _ := io.ReadAll(resp.Body)
		return respBody, fmt.Errorf("resp.Status: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll(resp.Body): %w", err)
	}

	return respBody, nil
}
