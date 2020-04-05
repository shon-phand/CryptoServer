package domain

import (
	"github.com/shon-phand/CryptoServer/datasources/redis"
	"github.com/shon-phand/CryptoServer/utils/errors"
)

func (cr *Currency) Get(cur string) (*Currency, *errors.RestErr) {

	res, err := redis.Client.HGetAll("sym:" + cur).Result()
	if err != nil {
		return nil, errors.StatusInternalServerError("database error")
	}

	_, ok := res["id"]
	if !ok {
		return nil, errors.StatusNotFoundError("currency not found")
	}

	Response := Currency{}
	Response.ID = res["id"]
	Response.Ask = res["ask"]
	Response.Bid = res["bid"]
	Response.Last = res["last"]
	Response.Open = res["open"]
	Response.Low = res["low"]
	Response.High = res["high"]
	Response.FeeCurrency = res["feeCurrency"]

	return &Response, nil
}

func (cr *Currencies) GetAll() (*Currencies, *errors.RestErr) {

	hashes, err := redis.Client.Keys("sym:*").Result()

	if err != nil {
		return nil, errors.StatusInternalServerError("database error")
	}

	if len(hashes) == 0 {
		return nil, errors.StatusNotFoundError("no currencies found in the database, please sync all currencies first")
	}

	result := Currencies{}
	Response := Currency{}
	for _, v := range hashes {

		res, err := redis.Client.HGetAll(v).Result()
		if err != nil {
			return nil, errors.StatusInternalServerError("database error")
		}
		Response.ID = res["id"]
		Response.Ask = res["ask"]
		Response.Bid = res["bid"]
		Response.Last = res["last"]
		Response.Open = res["open"]
		Response.Low = res["low"]
		Response.High = res["high"]
		Response.FeeCurrency = res["feeCurrency"]
		result = append(result, Response)

	}

	return &result, nil
}
