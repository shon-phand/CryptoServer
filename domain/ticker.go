package domain

type Ticker struct {
	Symbol      string `json:"symbol"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	Volume      string `json:"volume"`
	VolumeQuote string `json:"volumeQuote"`
	Timestamp   string `json:"timestamp"`
}

type Tickers []Ticker

//var tickers Tickers

// func SyncTickers() {

// 	resp, err := http.Get("https://api.hitbtc.com/api/2/public/ticker")
// 	if err != nil {
// 		fmt.Println("error")
// 	}
// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	err = json.Unmarshal(body, &tickers)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	for _, v := range tickers {
// 		_, err := redis.Client.HSet("tic:"+v.Symbol, []string{"symbol", v.Symbol, "ask", v.Ask, "bid", v.Bid, "last", v.Last, "open", v.Open, "low", v.Low, "high", v.High, "volume", v.Volume, "volumeQuote", v.VolumeQuote}).Result()
// 		if err != nil {
// 			fmt.Println("error in saving tic")
// 		}

// 	}
// }
