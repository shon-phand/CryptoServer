package services

import (
	"fmt"

	"github.com/shon-phand/CryptoServer/datasources/redis"
	"github.com/shon-phand/CryptoServer/sync"
)

func Save(curr sync.Currencies) {

	for _, cr := range curr {
		ok, err := redis.Client.HSet("sym:"+cr.Symbol, "id", cr.ID, "ask", cr.Ask, "bid", cr.Bid, "last", cr.Last, "open", cr.Open, "low", cr.Low, "high", cr.High, "feeCurrency", cr.FeeCurrency).Result()
		if err != nil {
			fmt.Println("error in saving in redis")
		}
		fmt.Println("added currency", ok, " ", cr)
	}

}
