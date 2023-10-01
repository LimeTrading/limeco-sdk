package limecosdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"

	"nhooyr.io/websocket"
)

type LimeClient struct {
	httpClient *http.Client
	apiKey     string
}

func NewLimeClient(apiKey string, timeout time.Duration, httpClient *http.Client) (client *LimeClient) {
	client.apiKey = apiKey

	if httpClient == nil {
		client.httpClient = http.DefaultClient
	} else {
		client.httpClient = httpClient
	}

	client.httpClient.Timeout = timeout

	return
}

func httpDo[B, R any](client *LimeClient, method string, args url.Values, body B, endpoint ...string) (out R, err error) {
	u := url.URL{
		Scheme:   "https",
		Host:     "api.lime.co",
		Path:     path.Join(endpoint...),
		RawQuery: args.Encode(),
	}

	bits, err := json.Marshal(body)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", u.String(), bytes.NewReader(bits))
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accepts", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	bits, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(bits, &out)

	return
}

func connect[S ~string](client *LimeClient, endpoint string) (out *LiveMarketData[S], err error) {
	out.client = client

	ctx, cancel := context.WithTimeout(context.Background(), client.httpClient.Timeout)
	defer cancel()

	out.ws, _, err = websocket.Dial(ctx, "wss://api.lime.co/"+endpoint, &websocket.DialOptions{
		CompressionMode: websocket.CompressionContextTakeover,
		HTTPClient:      client.httpClient,
		HTTPHeader: http.Header{
			"Authorization": []string{fmt.Sprintf("Bearer %s", client.apiKey)},
		},
	})

	return
}

func (client *LimeClient) ConnectToMarketData() (out *LiveMarketData[MarketDataAction], err error) {
	return connect[MarketDataAction](client, "marketData")
}

func (client *LimeClient) ConnectToAccountsFeed() (out *LiveMarketData[AccountDataAction], err error) {
	return connect[AccountDataAction](client, "accounts")
}
