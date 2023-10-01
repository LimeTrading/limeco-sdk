package limecosdk

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// https://docs.lime.co/trader/accounts/
func (client *LimeClient) GetAccountBalances() (out chan []Account, e chan error) {
	out = make(chan []Account)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		a, err := httpDo[any, []Account](client, "GET", nil, nil, "accounts")
		if err != nil {
			e <- err
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/accounts/get-account-positions
func (client *LimeClient) GetAccountPositions(accountNumber uint, date time.Time) (out chan []StockPosition, e chan error) {
	out = make(chan []StockPosition)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		u := url.Values{
			"date": []string{date.String()},
		}

		// a, err := httpDo[[]StockPosition](client, "GET", u, nil, "accounts", fmt.Sprintf("%v", accountNumber))
		a, err := httpDo[any, []StockPosition](client, "GET", u, nil, "accounts", fmt.Sprintf("%v", accountNumber))

		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/accounts/get-account-trades
func (client *LimeClient) GetAccountTrades(accountNumber uint, date time.Time, limit, skip uint) (out chan []StockPosition, e chan error) {
	out = make(chan []StockPosition)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		u := url.Values{
			"limit": []string{fmt.Sprintf("%v", limit)},
			"skip":  []string{fmt.Sprintf("%v", skip)},
		}

		// a, err := httpDo[[]StockPosition](client, "GET", u, nil, "accounts", fmt.Sprintf("%v", accountNumber))
		a, err := httpDo[any, []StockPosition](client, "GET", u, nil, "accounts", fmt.Sprintf("%v", accountNumber))

		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/accounts/get-account-trades
func (client *LimeClient) GetTransactionJournal(accountNumber uint, start, end time.Time, limit, skip uint) (out chan TransactionsJournal, e chan error) {
	out = make(chan TransactionsJournal)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		u := url.Values{
			"startDate": []string{start.String()},
			"endDate":   []string{end.String()},
			"limit":     []string{fmt.Sprintf("%v", limit)},
			"skip":      []string{fmt.Sprintf("%v", skip)},
		}

		a, err := httpDo[any, TransactionsJournal](client, "GET", u, nil, "accounts", fmt.Sprintf("%v", accountNumber), "transactions")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/trading/
func (client *LimeClient) PlaceOrder(order Order) (out chan OrderStatus, e chan error) {
	out = make(chan OrderStatus)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		missingField := isStructFieldsSet(order)
		if missingField != "" {
			e <- errors.New(fmt.Sprintf("%s field not set", missingField))
			return
		}

		a, err := httpDo[Order, OrderStatus](client, "POST", nil, order, "orders", "place")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/trading/validate-order
func (client *LimeClient) ValidateOrder(order Order) (out chan ValidationStatus, e chan error) {
	out = make(chan ValidationStatus)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		missingField := isStructFieldsSet(order)
		if missingField != "" {
			e <- errors.New(fmt.Sprintf("%s field not set", missingField))
			return
		}

		// a, err := httpPost[Order, ValidationStatus](client, order, "orders", "validate")
		a, err := httpDo[Order, ValidationStatus](client, "POST", nil, order, "orders", "validate")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/trading/get-order-details
func (client *LimeClient) GetOrderDetails(orderId string) (out chan OrderDetails, e chan error) {
	out = make(chan OrderDetails)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		a, err := httpDo[any, OrderDetails](client, "GET", nil, nil, "orders", orderId)
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/trading/get-active-orders
func (client *LimeClient) GetActiveOrders(accountNumber uint) (out chan []OrderDetails, e chan error) {
	out = make(chan []OrderDetails)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		a, err := httpDo[any, []OrderDetails](client, "GET", nil, nil, "accounts", fmt.Sprintf("%v", accountNumber), "activeOrders")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/trading/cancel-order
func (client *LimeClient) CancelOrder(orderId uint) (out chan OrderStatus, e chan error) {
	out = make(chan OrderStatus)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		a, err := httpDo[any, OrderStatus](client, "GET", nil, nil, "accounts", fmt.Sprintf("%v", orderId), "activeOrders")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/trading/estimate-fee-charges
func (client *LimeClient) EstimateFees(order Order) (out chan []EstimatedFee, e chan error) {
	out = make(chan []EstimatedFee)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		a, err := httpDo[Order, []EstimatedFee](client, "GET", nil, order, "pricing", "fees")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/market-data/get-current-quote-array
func (client *LimeClient) GetQuotes(symbols []string) (out chan []Quote, e chan error) {
	out = make(chan []Quote)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		a, err := httpDo[[]string, []Quote](client, "GET", nil, symbols, "pricing", "fees")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/market-data/get-quotes-history
func (client *LimeClient) GetQuoteHistory(symbol, period string, start, end time.Time) (out chan []Quote, e chan error) {
	out = make(chan []Quote)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		u := url.Values{
			"symbol": []string{symbol},
			"period": []string{period},
			"from":   []string{start.Format(time.RFC3339)},
			"to":     []string{end.Format(time.RFC3339)},
		}

		a, err := httpDo[any, []Quote](client, "GET", u, nil, "pricing", "fees")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/market-data/get-trading-schedule
func (client *LimeClient) GetSchedule() (out chan Schedule, e chan error) {
	out = make(chan Schedule)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		a, err := httpDo[any, Schedule](client, "GET", nil, nil, "schedule")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/market-data/lookup-securities
func (client *LimeClient) LookupSecurities(partialSymbol string, limit uint) (out chan SymbolLookupResults, e chan error) {
	out = make(chan SymbolLookupResults)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		u := url.Values{
			"query": []string{partialSymbol},
			"limit": []string{fmt.Sprintf("%v", limit)},
		}

		a, err := httpDo[any, SymbolLookupResults](client, "GET", u, nil, "schedule")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/market-data/get-option-series
func (client *LimeClient) GetOptionSeries(symbol string) (out chan []OptionSeries, e chan error) {
	out = make(chan []OptionSeries)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		a, err := httpDo[any, []OptionSeries](client, "GET", nil, nil, "securities", symbol, "options", "series")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}

// https://docs.lime.co/trader/market-data/get-option-chain
func (client *LimeClient) GetOptionChain(symbol, series string, expiration time.Time) (out chan OptionSeriesChain, e chan error) {
	out = make(chan OptionSeriesChain)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		u := url.Values{
			"expiration": []string{expiration.Format(time.RFC3339)},
			"series":     []string{series},
		}

		a, err := httpDo[any, OptionSeriesChain](client, "GET", u, nil, "securities", symbol, "options")
		if err != nil {
			e <- err
			return
		}

		out <- a
	}()

	return
}
