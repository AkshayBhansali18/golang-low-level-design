package order_worker

import (
	"github.com/rs/zerolog/log"
	"stockbroker/models"
	"stockbroker/orderexecutor"
	"stockbroker/user"
	"stockbroker/validation"
)

type OrderWorker interface {
	ExecuteOrder(userInfo *user.UserInfo, orderInfo *models.Order) error
}

type orderWorkerImpl struct {
	validatorService validation.OrderValidator
	orderImplementor orderexecutor.OrderExecutor
}

func NewOrderWorker(validatorService validation.OrderValidator, orderImplementor orderexecutor.OrderExecutor) OrderWorker {
	return &orderWorkerImpl{
		validatorService: validatorService,
		orderImplementor: orderImplementor,
	}
}

func (o *orderWorkerImpl) ExecuteOrder(userInfo *user.UserInfo, orderInfo *models.Order) error {
	err := o.validatorService.ValidateOrder(userInfo, orderInfo)
	if err != nil {
		log.Error().Msgf("error in validating order: %v", err)
		return err
	}
	err = o.orderImplementor.PlaceOrder()
	if err != nil {
		log.Error().Msgf("error in placing order: %v", err)
		return err
	}
	return nil
}
