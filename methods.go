package limecosdk

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

// https://docs.lime.co/trader/accounts/
func (client *LimeClient) GetAccountBalances() (out []Account, err error) {
	out, err = httpDo[any, []Account](client, "GET", nil, nil, "accounts")
	return
}

// https://docs.lime.co/trader/accounts/get-account-positions
func (client *LimeClient) GetAccountPositions(accountNumber uint, date time.Time) (out []StockPosition, err error) {
	u := url.Values{
		"date": []string{date.String()},
	}

	out, err = httpDo[any, []StockPosition](client, "GET", u, nil, "accounts", fmt.Sprintf("%v", accountNumber))
	return
}

// https://docs.lime.co/trader/accounts/get-account-trades
func (client *LimeClient) GetAccountTrades(accountNumber uint, date time.Time, limit, skip uint) (out []StockPosition, err error) {
	u := url.Values{
		"limit": []string{fmt.Sprintf("%v", limit)},
		"skip":  []string{fmt.Sprintf("%v", skip)},
	}

	out, err = httpDo[any, []StockPosition](client, "GET", u, nil, "accounts", fmt.Sprintf("%v", accountNumber))
	return
}

// https://docs.lime.co/trader/accounts/transactions-journal
func (client *LimeClient) GetTransactionJournal(accountNumber uint, start, end time.Time, limit, skip uint) (out TransactionsJournal, err error) {
	u := url.Values{
		"startDate": []string{start.String()},
		"endDate":   []string{end.String()},
		"limit":     []string{fmt.Sprintf("%v", limit)},
		"skip":      []string{fmt.Sprintf("%v", skip)},
	}

	out, err = httpDo[any, TransactionsJournal](client, "GET", u, nil, "accounts", fmt.Sprintf("%v", accountNumber), "transactions")

	return
}

// https://docs.lime.co/trader/trading/
func (client *LimeClient) PlaceOrder(order Order) (out OrderStatus, err error) {
	missingField := isStructFieldsSet(order)
	if missingField != "" {
		err = errors.New(fmt.Sprintf("%s field not set", missingField))
		return
	}

	out, err = httpDo[Order, OrderStatus](client, "POST", nil, order, "orders", "place")

	return
}

// https://docs.lime.co/trader/trading/validate-order
func (client *LimeClient) ValidateOrder(order Order) (out ValidationStatus, err error) {
	missingField := isStructFieldsSet(order)
	if missingField != "" {
		err = errors.New(fmt.Sprintf("%s field not set", missingField))
		return
	}

	out, err = httpDo[Order, ValidationStatus](client, "POST", nil, order, "orders", "validate")
	return
}

// https://docs.lime.co/trader/trading/get-order-details
func (client *LimeClient) GetOrderDetails(orderId string) (out OrderDetails, err error) {
	out, err = httpDo[any, OrderDetails](client, "GET", nil, nil, "orders", orderId)
	return
}

// https://docs.lime.co/trader/trading/get-active-orders
func (client *LimeClient) GetActiveOrders(accountNumber uint) (out []OrderDetails, err error) {
	out, err = httpDo[any, []OrderDetails](client, "GET", nil, nil, "accounts", fmt.Sprintf("%v", accountNumber), "activeOrders")
	return
}

// https://docs.lime.co/trader/trading/cancel-order
func (client *LimeClient) CancelOrder(orderId uint) (out OrderStatus, err error) {
	out, err = httpDo[any, OrderStatus](client, "GET", nil, nil, "accounts", fmt.Sprintf("%v", orderId), "activeOrders")
	return
}

// https://docs.lime.co/trader/trading/estimate-fee-charges
func (client *LimeClient) EstimateFees(order Order) (out []EstimatedFee, err error) {
	out, err = httpDo[Order, []EstimatedFee](client, "GET", nil, order, "pricing", "fees")
	return
}

// https://docs.lime.co/trader/market-data/get-current-quote-array
func (client *LimeClient) GetQuotes(symbols []string) (out []Quote, err error) {
	out, err = httpDo[[]string, []Quote](client, "GET", nil, symbols, "pricing", "fees")
	return
}

// https://docs.lime.co/trader/market-data/get-quotes-history
func (client *LimeClient) GetQuoteHistory(symbol, period string, start, end time.Time) (out []Quote, err error) {
	u := url.Values{
		"symbol": []string{symbol},
		"period": []string{period},
		"from":   []string{start.Format(time.RFC3339)},
		"to":     []string{end.Format(time.RFC3339)},
	}

	out, err = httpDo[any, []Quote](client, "GET", u, nil, "pricing", "fees")
	return
}

// https://docs.lime.co/trader/market-data/get-trading-schedule
func (client *LimeClient) GetSchedule() (out Schedule, err error) {
	out, err = httpDo[any, Schedule](client, "GET", nil, nil, "schedule")
	return
}

// https://docs.lime.co/trader/market-data/lookup-securities
func (client *LimeClient) LookupSecurities(partialSymbol string, limit uint) (out SymbolLookupResults, err error) {
	u := url.Values{
		"query": []string{partialSymbol},
		"limit": []string{fmt.Sprintf("%v", limit)},
	}

	out, err = httpDo[any, SymbolLookupResults](client, "GET", u, nil, "schedule")
	return
}

// https://docs.lime.co/trader/market-data/get-option-series
func (client *LimeClient) GetOptionSeries(symbol string) (out []OptionSeries, err error) {
	out, err = httpDo[any, []OptionSeries](client, "GET", nil, nil, "securities", symbol, "options", "series")
	return
}

// https://docs.lime.co/trader/market-data/get-option-chain
func (client *LimeClient) GetOptionChain(symbol, series string, expiration time.Time) (out OptionSeriesChain, err error) {
	u := url.Values{
		"expiration": []string{expiration.Format(time.RFC3339)},
		"series":     []string{series},
	}

	out, err = httpDo[any, OptionSeriesChain](client, "GET", u, nil, "securities", symbol, "options")
	return
}
