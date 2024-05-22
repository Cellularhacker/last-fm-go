package lastFm

import (
	"crypto/tls"
	"net/http"
	"time"
)

var client *Client

func initClient() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 1000
	t.MaxIdleConnsPerHost = 1000
	t.IdleConnTimeout = 900 * time.Second
	t.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	client = &Client{
		client: &http.Client{
			Timeout:   600 * time.Second,
			Transport: t,
		},
		maxRetries: 5,
		retryDelay: 1 * time.Second,
	}
}

func GetClient() *Client {
	return client
}

type Client struct {
	client     *http.Client
	maxRetries int
	retryDelay time.Duration
}

func (dc *Client) SetMaxRetries(maxRetries int) *Client {
	nDc := dc
	// MARK: Safety for nil point error
	if nDc == nil {
		nDc = GetClient()
	}
	// MARK: Safety for minus integer numbers
	if maxRetries < 0 {
		maxRetries = 0
	}
	nDc.maxRetries = maxRetries

	return nDc
}

func (dc *Client) SetRetryDelay(retryDelay time.Duration) *Client {
	nDc := dc
	if nDc == nil {
		nDc = GetClient()
	}
	if retryDelay.Milliseconds() < 0 {
		retryDelay = 0
	}
	nDc.retryDelay = retryDelay

	return nDc
}

func (dc *Client) GetMaxRetriesCount() int {
	return dc.maxRetries
}

func (dc *Client) GetRetryDelay() time.Duration {
	return dc.retryDelay
}

func (dc *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := &http.Response{}, error(nil)

	for i := 0; i < dc.GetMaxRetriesCount(); i++ {
		resp, err = dc.client.Do(req)
		if err == nil {
			return resp, nil
		}

		<-time.NewTicker(dc.retryDelay).C
	}

	return resp, err
}
