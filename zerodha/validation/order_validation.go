package validation

import (
	"errors"
	"github.com/rs/zerolog/log"
	"stockbroker/models"
	"stockbroker/user"
)

type orderValidatorImpl struct {
}

type OrderValidator interface {
	ValidateOrder(userInfo *user.UserInfo, orderInfo *models.Order) error
}

func NewOrderValidation() OrderValidator {
	return &orderValidatorImpl{}
}

func (o *orderValidatorImpl) ValidateOrder(userInfo *user.UserInfo, orderInfo *models.Order) error {
	if orderInfo.TotalPrice > userInfo.FundsRemaining {
		return errors.New("insufficient funds")
	}
	log.Info().Msg("Order validated successfully!")
	return nil
}
