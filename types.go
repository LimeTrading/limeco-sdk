package limecosdk

import (
	"time"

	"nhooyr.io/websocket"
)

// All your information about an account
type Account struct {
	AccountNumber         string  `json:"account_number"`
	TradePlatform         string  `json:"trade_platform"`
	MarginType            string  `json:"margin_type"`
	Restriction           string  `json:"restriction"`
	DayTradesCount        int     `json:"daytrades_count"`
	AccountValueTotal     float64 `json:"account_value_total"`
	Cash                  float64 `json:"cash"`
	DayTradingBuyingPower float64 `json:"day_trading_buying_power"`
	MarginBuyingPower     float64 `json:"margin_buying_power"`
	NonMarginBuyingPower  float64 `json:"non_margin_buying_power"`
	PositionMarketValue   float64 `json:"position_market_value"`
	UnsettledCash         float64 `json:"unsettled_cash"`
	CashToWithdraw        float64 `json:"cash_to_withdraw"`
}

// Your position with a stock
type StockPosition struct {
	Symbol           string  `json:"symbol"`
	Quantity         int     `json:"quantity"`
	AverageOpenPrice float64 `json:"average_open_price"`
	CurrentPrice     float64 `json:"current_price"`
	SecurityType     string  `json:"security_type"`
}

// Represents a trade
type Trade struct {
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"timestamp"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Side      string  `json:"side"`
}

// Represents a transaction
type Transaction struct {
	ID          string           `json:"id"`
	Type        string           `json:"type"`
	Description string           `json:"description"`
	Date        string           `json:"date"`
	Asset       Asset            `json:"asset,omitempty"`
	Cash        Cash             `json:"cash"`
	Fees        []TransactionFee `json:"fees"`
}

// Represents an asset you own
type Asset struct {
	Symbol            string  `json:"symbol"`
	SymbolDescription string  `json:"symbol_description"`
	Quantity          int     `json:"quantity"`
	Price             float64 `json:"price"`
}

// Cash amounts
type Cash struct {
	GrossAmount float64 `json:"gross_amount"`
	NetAmount   float64 `json:"net_amount"`
}

// Transaction fee information
type TransactionFee struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

// Information about your transactions
type TransactionData struct {
	Transactions []Transaction `json:"transactions"`
	Count        int           `json:"count"`
}

// Information about your trades
type Trades struct {
	Trades []Trade `json:"trades"`
	Count  uint    `json:"count"`
}

// A journal record of your transactions
type TransactionsJournal struct {
	Transactions []Transaction `json:"transactions"`
	Count        uint          `json:"count"`
}

// Status of an order
type OrderStatus struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

type TimeInForceValue string

const TIF_Day TimeInForceValue = "day"
const TIF_AfterMarket TimeInForceValue = "ext"

type MarketOrderType string

const MarketOrderType_Market MarketOrderType = "market"
const MarketOrderType_Limit MarketOrderType = "limit"

type OrderSide string

const OrderSide_Buy OrderSide = "buy"
const OrderSide_Sell OrderSide = "sell"

// Represents an order to submit
type Order struct {
	AccountNumber string           `json:"account_number"`
	Symbol        string           `json:"symbol"`
	Quantity      int              `json:"quantity"`
	Price         float64          `json:"price"`
	TimeInForce   TimeInForceValue `json:"time_in_force"`
	OrderType     MarketOrderType  `json:"order_type"`
	Side          OrderSide        `json:"side"`
	Exchange      string           `json:"exchange"`
}

// Is action valid
type ValidationStatus struct {
	IsValid bool   `json:"is_valid"`
	Message string `json:"validation_message"`
}

// Details about an order
type OrderDetails struct {
	AccountNumber string  `json:"account_number"`
	Symbol        string  `json:"symbol"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	TimeInForce   string  `json:"time_in_force"`
	OrderType     string  `json:"order_type"`
	Side          string  `json:"side"`
	Exchange      string  `json:"exchange"`
}

// Fee estimation
type EstimatedFee struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

// Information about a quote
type Quote struct {
	Symbol            string  `json:"symbol"`
	Ask               float64 `json:"ask"`
	AskSize           int     `json:"ask_size"`
	Bid               float64 `json:"bid"`
	BidSize           int     `json:"bid_size"`
	Last              float64 `json:"last"`
	LastSize          int     `json:"last_size"`
	Volume            int     `json:"volume"`
	Date              int64   `json:"date"`
	High              float64 `json:"high"`
	Low               float64 `json:"low"`
	Open              float64 `json:"open"`
	Close             float64 `json:"close"`
	Week52High        float64 `json:"week52_high"`
	Week52Low         float64 `json:"week52_low"`
	Change            float64 `json:"change"`
	ChangePercentage  float64 `json:"change_pc"`
	OpenInterest      int     `json:"open_interest"`
	ImpliedVolatility float64 `json:"implied_volatility"`
	TheoreticalPrice  float64 `json:"theoretical_price"`
	Delta             float64 `json:"delta"`
	Gamma             float64 `json:"gamma"`
	Theta             float64 `json:"theta"`
	Vega              float64 `json:"vega"`
}

// When you're allowed to trade
type Schedule struct {
	Session string `json:"session"`
}

// Information about a symbol
type Symbol struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

// Results of a symbol lookup
type SymbolLookupResults struct {
	Securities []Symbol `json:"securities"`
	Count      uint     `json:"count"`
}

// Information about an options series
type OptionSeries struct {
	Series      string      `json:"series"`
	Expirations []time.Time `json:"expirations"`
	Size        uint        `json:"contract_size"`
}

// An item in a chain of options
type OptionSeriesChainItem struct {
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	Strike uint   `json:"strike"`
}

// Information about your options
type OptionSeriesChain struct {
	Size  uint                    `json:"contract_size"`
	Style string                  `json:"style"`
	Chain []OptionSeriesChainItem `json:"chain"`
}

// A directive about an account
type Directive struct {
	Action  string `json:"action"`
	Account string `json:"account"`
}

// Information about a position
type PositionDataItem struct {
	Symbol           string  `json:"symbol"`
	AverageOpenPrice float64 `json:"average_open_price"`
	CurrentPrice     float64 `json:"current_price"`
	Quantity         int     `json:"quantity"`
	SecurityType     string  `json:"security_type"`
}

type PositionData struct {
	Account   string             `json:"account"`
	Positions []PositionDataItem `json:"positions"`
}

// Information about your balance
type BalanceData struct {
	AccountNumber        uint    `json:"account_number"`
	TradePlatform        string  `json:"trade_platform"`
	MarginType           string  `json:"margin_type"`
	Restriction          string  `json:"restriction"`
	AccountValueTotal    float64 `json:"account_value_total"`
	Cash                 float64 `json:"cash"`
	MarginBuyingPower    float64 `json:"margin_buying_power"`
	NonMarginBuyingPower float64 `json:"non_margin_buying_power"`
	PositionMarketValue  float64 `json:"position_market_value"`
	UnsettledCash        float64 `json:"unsettled_cash"`
	CashToWithdraw       float64 `json:"cash_to_withdraw"`
}

// Information about an order
type OrderData struct {
	AccountNumber     string  `json:"account_number"`
	ClientID          string  `json:"client_id"`
	Exchange          string  `json:"exchange"`
	Quantity          int     `json:"quantity"`
	ExecutedQuantity  int     `json:"executed_quantity"`
	OrderStatus       string  `json:"order_status"`
	Price             float64 `json:"price"`
	StopPrice         float64 `json:"stop_price"`
	TimeInForce       string  `json:"time_in_force"`
	OrderType         string  `json:"order_type"`
	OrderSide         string  `json:"order_side"`
	Symbol            string  `json:"symbol"`
	ExecutedPrice     float64 `json:"executed_price"`
	Comment           string  `json:"comment"`
	ExecutedTimestamp int64   `json:"executed_timestamp"`
}

// Information about a trade
type TradeData struct {
	AccountNumber string  `json:"account_number"`
	Symbol        string  `json:"symbol"`
	Timestamp     int64   `json:"timestamp"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Amount        float64 `json:"amount"`
	Side          string  `json:"side"`
}

// A connection to the market
type LiveMarketData[ActionType any] struct {
	client *LimeClient
	ws     *websocket.Conn
}

// Market data update
type MarketData struct {
	Type             string  `json:"t"`
	Symbol           string  `json:"s"`
	LastSize         int     `json:"ls"`
	LastMarket       string  `json:"lm"`
	Ask              float64 `json:"a,omitempty"`
	AskSize          int     `json:"as,omitempty"`
	Bid              float64 `json:"b,omitempty"`
	BidSize          int     `json:"bs,omitempty"`
	Last             float64 `json:"l,omitempty"`
	High             float64 `json:"high,omitempty"`
	Low              float64 `json:"low,omitempty"`
	Open             float64 `json:"open,omitempty"`
	Close            float64 `json:"close,omitempty"`
	Change           float64 `json:"change,omitempty"`
	ChangePercentage float64 `json:"change_pc,omitempty"`
	Timestamp        int64   `json:"d,omitempty"`
	Volume           int     `json:"v,omitempty"`
}

// Sends an action on an account data connection
type AccountDataActionCommand struct {
	Action  AccountDataAction `json:"action"`
	Account string            `json:"account"`
}

// Send an action on a market data connection
type MarketDataActionCommand struct {
	Action  MarketDataAction `json:"action"`
	Symbols []string         `json:"symbols"`
}

type rawAccountData struct {
	Type string     `json:"t"`
	Data []struct{} `json:"data"`
}
