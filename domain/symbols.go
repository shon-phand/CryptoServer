package domain

type Symbol struct {
	ID                   string `json:"id"`
	BaseCurrency         string `json:"baseCurrency"`
	QuoteCurrency        string `json:"quoteCurrency"`
	QuantityIncrement    string `json:"quantityIncrement"`
	TickSize             string `json:"tickSize"`
	TakeLiquidityRate    string `json:"takeLiquidityRate"`
	ProvideLiquidityRate string `json:"provideLiquidityRate"`
	FeeCurrency          string `json:"feeCurrency"`
}

type Symbols []Symbol

// type Response struct {
// 	Data []Symbol `json:"Data"`
// }

// var symbols Symbols

// func SyncSymbols() {

// 	resp, err := http.Get("https://api.hitbtc.com/api/2/public/symbol")
// 	if err != nil {
// 		fmt.Println("error")
// 	}
// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	err = json.Unmarshal(body, &symbols)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	for _, v := range symbols {
// 		_, err := redis.Client.HSet("sym:"+v.ID, []string{"id", v.ID, "baseCurrency", v.BaseCurrency, "quoteCurrency", v.QuoteCurrency, "quantityIncrement", v.QuantityIncrement, "tickSize", v.TickSize, "takeLiquidityRate", v.TakeLiquidityRate, "provideLiquidityRate", v.ProvideLiquidityRate, "feeCurrency", v.FeeCurrency}).Result()
// 		if err != nil {
// 			fmt.Println("error in saving symbols")
// 		}
// 	}
// }
