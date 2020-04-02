package domain

type Currency struct {
	ID          string `json:"id"`
	Ask         string `json:"ask"`
	Bid         string `json:"bid"`
	Last        string `json:"last"`
	Open        string `json:"open"`
	Low         string `json:"low"`
	High        string `json:"high"`
	FeeCurrency string `json:"feeCurrency"`
}

type Currencies []Currency

// var currencies Currencies

// func SyncCurrencies() {
// 	resp, err := http.Get("https://api.hitbtc.com/api/2/public/currency")
// 	if err != nil {
// 		fmt.Println("error")
// 	}
// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	err = json.Unmarshal(body, &currencies)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	for _, v := range currencies {
// 		_, err := redis.Client.HSet("curr:"+v.ID, "id", v.ID, "fullName", v.FullName, "crypto", v.Crypto, "payinEnabled", v.PayinEnabled, "payinConfirmations", v.PayinConfirmations, "payinPaymentId", v.PayinPaymentID, "payoutEnabled", v.PayoutEnabled, "ayoutIsPaymentID", v.PayoutIsPaymentID, "transferEnabled", v.TransferEnabled, "delisted", v.Delisted, "payoutFee", v.PayoutFee).Result()
// 		if err != nil {
// 			fmt.Println("error in saving currencies")
// 		}
// 	}
// }
