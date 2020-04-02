package sync

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/shon-phand/CryptoServer/domain"
)

type Crnc struct {
	Symbol      string `json:"symbol"`
	ID          string `json:"id"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"feeCurrency"`
}

type Currencies []Crnc

var AllCurrencies Currencies
var Response Crnc

func SyncCurrency() Currencies {

	type Symbols []domain.Symbol
	type Tickers domain.Ticker

	rawSymbols, err := http.Get("https://api.hitbtc.com/api/2/public/symbol")
	if err != nil {
		fmt.Println("error in downloading the symbols")
	}
	defer rawSymbols.Body.Close()
	var symbols Symbols
	body, _ := ioutil.ReadAll(rawSymbols.Body)
	err = json.Unmarshal(body, &symbols)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range symbols {

		rawTckr, err := http.Get("https://api.hitbtc.com/api/2/public/ticker/" + v.ID)
		if err != nil {
			fmt.Println("error in fetching ticker")
		}
		defer rawTckr.Body.Close()

		var ticekrs Tickers

		body, _ := ioutil.ReadAll(rawTckr.Body)
		err = json.Unmarshal(body, &ticekrs)
		if err != nil {
			fmt.Println(err.Error())
		}
		Response.Symbol = v.ID
		Response.ID = v.BaseCurrency
		Response.Ask = ticekrs.Ask
		Response.Bid = ticekrs.Bid
		Response.Last = ticekrs.Last
		Response.Open = ticekrs.Open
		Response.Low = ticekrs.Low
		Response.High = ticekrs.High
		Response.FeeCurrency = v.FeeCurrency
		AllCurrencies = append(AllCurrencies, Response)
		fmt.Println("added currency", Response)

	}

	return AllCurrencies
}
