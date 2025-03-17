package orderexecutor

import (
	"github.com/rs/zerolog/log"
	"stockbroker/clients"
)

type OrderExecutor interface {
	PlaceOrder() error
}

type orderExecutorImpl struct {
	clients.StockExchangeClient
}

func NewOrderImplementor(stockExchangeClient clients.StockExchangeClient) OrderExecutor {
	return &orderExecutorImpl{
		StockExchangeClient: stockExchangeClient,
	}
}

func (o *orderExecutorImpl) PlaceOrder() error {
	err := o.StockExchangeClient.EvaluateStockPrice()
	if err != nil {
		log.Error().Msgf("error in validating order: %v", err)
		return err
	}
	log.Info().Msg("Order placed successfully")
	return nil
}
