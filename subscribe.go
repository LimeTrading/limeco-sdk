package limecosdk

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"nhooyr.io/websocket"
)

func listen[S any, T any](lmd *LiveMarketData[S]) (in chan S, out chan T, e chan error) {
	in = make(chan S)
	out = make(chan T)
	e = make(chan error)

	go func() {
		defer close(out)
		defer close(e)

		for s := range in {
			bits, err := json.Marshal(s)
			if err != nil {
				panic(err)
			}

			lmd.ws.Write(context.Background(), websocket.MessageText, bits)
		}
	}()

	go func() {
		defer close(out)
		defer close(e)

		for {
			msgType, bits, err := lmd.ws.Read(context.Background())
			if errors.Is(err, io.EOF) {
				return
			}
			if err != nil {
				e <- err
				continue
			}

			if msgType == websocket.MessageText {
				var m T
				err = json.Unmarshal(bits, &m)
				if err != nil {
					e <- err
					continue
				}

				out <- m
			} else {
				e <- errors.New("Unable to parse binary packets")
			}
		}
	}()

	return
}

func (lmd *LiveMarketData[MarketDataActionCommand]) ListenToMarket() (in chan MarketDataActionCommand, out chan MarketData, e chan error) {
	return listen[MarketDataActionCommand, MarketData](lmd)
}
