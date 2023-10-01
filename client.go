package limecosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
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
