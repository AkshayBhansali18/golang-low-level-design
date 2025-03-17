package main

import (
	"github.com/rs/zerolog/log"
	"stockbroker/clients"
	"stockbroker/models"
	"stockbroker/order_worker"
	"stockbroker/orderexecutor"
	"stockbroker/user"
	"stockbroker/validation"
)

func main() {
	log.Info().Msg("Hello, Welcome to Zerodha!")
	user := user.NewUser("Patrakar Popatlal", 100)

	orderInfo := &models.Order{
		TotalPrice:  50,
		Quantity:    10,
		OrderId:     "1",
		OrderType:   models.MARKET,
		OrderStatus: models.PENDING,
		Exchange:    models.NSE,
	}

	orderValidator := validation.NewOrderValidation()
	stockExchangeClient := clients.GetStockExchangeClient()
	orderExecutor := orderexecutor.NewOrderImplementor(stockExchangeClient)
	orderWorker := order_worker.NewOrderWorker(orderValidator, orderExecutor)
	err := orderWorker.ExecuteOrder(user, orderInfo)
	if err != nil {
		log.Error().Msgf("error in executing order: %v", err)
		return
	}
	log.Info().Msg("Order executed successfully!")
}
