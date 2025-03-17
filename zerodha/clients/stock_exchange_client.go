package clients

import (
	"github.com/rs/zerolog/log"
	"sync"
)

type StockExchangeClient interface {
	EvaluateStockPrice() error
}

type stockExchangeClientImpl struct{}

var stockExchangeClientInit sync.Once
var stockExchangeClientInstance *stockExchangeClientImpl

func GetStockExchangeClient() StockExchangeClient {
	//implementing singleton pattern here, to initialise the stockExchangeClient just once, and reuse the instance on every invocation
	stockExchangeClientInit.Do(func() {
		stockExchangeClientInstance = &stockExchangeClientImpl{}
	})
	return stockExchangeClientInstance
}

func (s *stockExchangeClientImpl) EvaluateStockPrice() error {
	//assuming successful API invocation
	log.Info().Msg("Stock price evaluated successfully!")
	return nil
}
