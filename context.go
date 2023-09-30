package limecosdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

type LimeClient struct {
	httpClient http.Client
	apiKey     string
}

func NewLimeClient(apiKey string, timeout time.Duration, httpClient *http.Client) (client *LimeClient) {
	client.apiKey = apiKey

	if httpClient == nil {
		client.httpClient = *http.DefaultClient
	} else {
		client.httpClient = *httpClient
	}

	client.httpClient.Timeout = timeout

	return
}

func (client *LimeClient) do(method string, args url.Values, endpoint ...string) ([]byte, error) {
	u := url.URL{
		Scheme:   "https",
		Host:     "api.lime.co",
		Path:     path.Join(endpoint...),
		RawQuery: args.Encode(),
	}
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", client.apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accepts", "application/json")

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func (client *LimeClient) GetAccountBalances() (out chan []Account, e chan error) {
	out = make(chan []Account)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		bytes, err := client.do("GET", nil, "accounts")
		if err != nil {
			e <- err
			return
		}

		var a []Account

		err = json.Unmarshal(bytes, &a)
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

func (client *LimeClient) GetAccountPositions(accountNumber uint, date time.Time) (out chan []StockPosition, e chan error) {
	out = make(chan []StockPosition)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		bytes, err := client.do("GET", url.Values{
			"date": []string{date.String()},
		}, "accounts", fmt.Sprintf("%v", accountNumber))
		if err != nil {
			e <- err
			return
		}

		var a []StockPosition

		err = json.Unmarshal(bytes, &a)
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

func (client *LimeClient) GetTransactionJournal(accountNumber uint, start, end time.Time, limit, skip uint) (out chan TransactionsJournal, e chan error) {
	out = make(chan TransactionsJournal)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		bytes, err := client.do("GET", url.Values{
			"startDate": []string{start.String()},
			"endDate":   []string{end.String()},
			"limit":     []string{fmt.Sprintf("%v", limit)},
			"skip":      []string{fmt.Sprintf("%v", skip)},
		}, "accounts", fmt.Sprintf("%v", accountNumber), "transactions")
		if err != nil {
			e <- err
			return
		}

		var a TransactionsJournal

		err = json.Unmarshal(bytes, &a)
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}
