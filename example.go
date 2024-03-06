package limecosdk

import (
	"log"
	"time"
)

func main() {
	const apiKey = ""

	client := NewLimeClient(apiKey, 5*time.Second, nil)

	marketData, err := client.ConnectToMarketData()
	if err != nil {
		panic(err)
	}

	actions, data, errs := marketData.ListenToMarket()

	actions <- MarketDataActionCommand{
		Action:  Subscribe,
		Symbols: []string{"F", "TSLA", "META", "GOOG"},
	}

	accountBalances, err := client.GetAccountBalances()
	if err != nil {
		panic(err)
	}

	myFavoriteAccount := accountBalances[0]

	for actions != nil && errs != nil {
		select {
		case info, ok := <-data:
			if !ok {
				data = nil
			}

			if info.Change > 1 {
				order := Order{
					Symbol:        info.Symbol,
					Quantity:      info.Volume * 2,
					AccountNumber: myFavoriteAccount.AccountNumber,
					Price:         info.Bid + 0.5,
					TimeInForce:   TIF_Day,
					OrderType:     MarketOrderType_Limit,
					Side:          OrderSide_Buy,
					Exchange:      "auto",
				}

				fees, err := client.EstimateFees(order)
				if err != nil {
					panic(err)
				}

				if fees[0].Amount > 1000 {
					panic("TOO EXPENSIVE!!!")
				}

				status, err := client.PlaceOrder(order)
				if err != nil {
					panic(err)
				}
				log.Println("Bought something", info.Symbol, status.Success)
			}
			break
		case err, ok := <-errs:
			if !ok {
				errs = nil
			}
			log.Println("Market data error:", err.Error())
			break
		}
	}
}
