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
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Asset       Asset  `json:"asset,omitempty"`
	Cash        Cash   `json:"cash"`
	Fees        []Fee  `json:"fees"`
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

type Fee struct {
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
