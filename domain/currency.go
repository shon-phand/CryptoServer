package domain

type Currency struct {
	ID                 string  `json:"id"`
	FullName           string  `json:"fullname"`
	Crypto             bool    `json:"crypto"`
	PayinEnabled       bool    `json:"payinenabled"`
	PayinPaymentID     bool    `json:"payinpaymentid"`
	PayinConfirmations int64   `json:"payinpayconfirmations"`
	PayoutEnabled      bool    `json:"payoutenabled"`
	PayoutIsPaymentID  bool    `json:"payoutispaymentid"`
	TransferEnabled    bool    `json:"transferenabled"`
	Delisted           bool    `json:"delisted"`
	PayoutFee          float64 `json:"payoutfee"`
}

type Currencies []Currency
