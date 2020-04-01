package domain

import (
	"strconv"

	"github.com/shon-phand/CryptoServer/datasources/redis"
	"github.com/shon-phand/CryptoServer/utils/errors"
)

func (cr *Currency) Get(cur string) (*Currency, *errors.RestErr) {

	res, err := redis.Client.HGetAll("curr:" + cur).Result()
	if err != nil {
		return nil, errors.StatusInternalServerError("database error")
	}

	curr := Currency{}
	curr.ID = res["id"]
	curr.FullName = res["fullName"]
	curr.Crypto, _ = strconv.ParseBool(res["crypto"])
	curr.PayinEnabled, _ = strconv.ParseBool(res["payinEnabled"])
	curr.PayinPaymentID, _ = strconv.ParseBool(res["payinPaymentId"])
	curr.PayinConfirmations, _ = strconv.ParseInt(res["payinConfirmations"], 10, 64)
	curr.PayoutEnabled, _ = strconv.ParseBool(res["payoutEnabled"])
	curr.PayoutIsPaymentID, _ = strconv.ParseBool(res["payoutIsPaymentId"])
	curr.TransferEnabled, _ = strconv.ParseBool(res["transferEnabled"])
	curr.Delisted, _ = strconv.ParseBool(res["delisted"])
	curr.PayoutFee, _ = strconv.ParseFloat(res["payoutFee"], 64)

	return &curr, nil
}

func (cr *Currencies) GetAll() (*Currencies, *errors.RestErr) {

	hashes, err := redis.Client.Keys("curr:*").Result()
	if err != nil {
		return nil, errors.StatusInternalServerError("database error")
	}
	result := Currencies{}
	for _, v := range hashes {

		curr := Currency{}
		res, err := redis.Client.HGetAll(v).Result()
		if err != nil {
			return nil, errors.StatusInternalServerError("database error")
		}
		curr.ID = res["id"]
		curr.FullName = res["fullName"]
		curr.Crypto, _ = strconv.ParseBool(res["crypto"])
		curr.PayinEnabled, _ = strconv.ParseBool(res["payinEnabled"])
		curr.PayinPaymentID, _ = strconv.ParseBool(res["payinPaymentId"])
		curr.PayinConfirmations, _ = strconv.ParseInt(res["payinConfirmations"], 10, 64)
		curr.PayoutEnabled, _ = strconv.ParseBool(res["payoutEnabled"])
		curr.PayoutIsPaymentID, _ = strconv.ParseBool(res["payoutIsPaymentId"])
		curr.TransferEnabled, _ = strconv.ParseBool(res["transferEnabled"])
		curr.Delisted, _ = strconv.ParseBool(res["delisted"])
		curr.PayoutFee, _ = strconv.ParseFloat(res["payoutFee"], 64)
		result = append(result, curr)

	}

	return &result, nil
}
