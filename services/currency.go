package services

import (
	"github.com/shon-phand/CryptoServer/domain"
	"github.com/shon-phand/CryptoServer/utils/errors"
)

type currencyService struct{}

var (
	CurrencyService currencyServiceInterface = &currencyService{}
)

type currencyServiceInterface interface {
	GetCurrency(string) (*domain.Currency, *errors.RestErr)
	GetAllCurrency() (*domain.Currencies, *errors.RestErr)
}

func (cs *currencyService) GetCurrency(curr string) (*domain.Currency, *errors.RestErr) {

	res := &domain.Currency{ID: curr}
	if err := res.Get(); err != nil {
		return nil, err
	}

	return res, nil

}

func (cs *currencyService) GetAllCurrency() (*domain.Currencies, *errors.RestErr) {

	res := &domain.Currency{}
	if err := res.GetAll(); err != nil {
		return nil, err
	}

	return res, nil

}
