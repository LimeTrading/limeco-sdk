package limecosdk

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

type StockPosition struct {
	Symbol           string  `json:"symbol"`
	Quantity         int     `json:"quantity"`
	AverageOpenPrice float64 `json:"average_open_price"`
	CurrentPrice     float64 `json:"current_price"`
	SecurityType     string  `json:"security_type"`
}

type Trade struct {
	Symbol    string  `json:"symbol"`
	Timestamp int64   `json:"timestamp"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Side      string  `json:"side"`
}

type Transaction struct {
	ID          string           `json:"id"`
	Type        string           `json:"type"`
	Description string           `json:"description"`
	Date        string           `json:"date"`
	Asset       Asset            `json:"asset,omitempty"`
	Cash        Cash             `json:"cash"`
	Fees        []TransactionFee `json:"fees"`
}

type Asset struct {
	Symbol            string  `json:"symbol"`
	SymbolDescription string  `json:"symbol_description"`
	Quantity          int     `json:"quantity"`
	Price             float64 `json:"price"`
}

type Cash struct {
	GrossAmount float64 `json:"gross_amount"`
	NetAmount   float64 `json:"net_amount"`
}

type TransactionFee struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type TransactionData struct {
	Transactions []Transaction `json:"transactions"`
	Count        int           `json:"count"`
}

type Trades struct {
	Trades []Trade `json:"trades"`
	Count  uint    `json:"count"`
}

type TransactionsJournal struct {
	Transactions []Transaction `json:"transactions"`
	Count        uint          `json:"count"`
}

type OrderStatus struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

type Order struct {
	AccountNumber string  `json:"account_number"`
	Symbol        string  `json:"symbol"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	TimeInForce   string  `json:"time_in_force"`
	OrderType     string  `json:"order_type"`
	Side          string  `json:"side"`
	Exchange      string  `json:"exchange"`
}

type ValidationStatus struct {
	IsValid bool   `json:"is_valid"`
	Message string `json:"validation_message"`
}

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

type EstimatedFee struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

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

type Schedule struct {
	Session string `json:"session"`
}

type Symbol struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

type SymbolLookupResults struct {
	Securities []Symbol `json:"securities"`
	Count      uint     `json:"count"`
}

type OptionSeries struct {
	Series      string      `json:"series"`
	Expirations []time.Time `json:"expirations"`
	Size        uint        `json:"contract_size"`
}

type OptionSeriesChainItem struct {
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
	Strike uint   `json:"strike"`
}

type OptionSeriesChain struct {
	Size  uint                    `json:"contract_size"`
	Style string                  `json:"style"`
	Chain []OptionSeriesChainItem `json:"chain"`
}
