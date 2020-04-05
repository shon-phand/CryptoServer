package services

import (
	"net/http"
	"testing"

	"github.com/shon-phand/CryptoServer/domain"
)

func TestGetAllCurrency(t *testing.T) {

	var csi currencyService

	res, err := csi.GetAllCurrency()

	if res == nil && err != nil {
		if err.Status == http.StatusInternalServerError {
			t.Errorf("database error")
		}
		if err.Status == http.StatusNotFound {
			t.Errorf("no currencis in database")
		}

	}

	if res != nil && err == nil {
		if len(*res) < 1 {
			t.Errorf("currencies not received")
		}
	}
}

func TestGetCurrency(t *testing.T) {

	var csi currencyService
	curr := "BTCUSD"
	res, err := csi.GetCurrency(curr)

	if res == nil && err != nil {
		if err.Status == http.StatusInternalServerError && err.Code == "service-error" && err.Message == "database error" {
			t.Errorf("database error")
		}
		if err.Status == http.StatusNotFound {
			t.Errorf("currencis not in present database")
		}

	}

	var test = domain.Currency{"BTC", "6779.37", "6779.03", "6778.12", "6862.59", "6644.30", "6996.27", "USD"}
	if res != nil && err == nil {
		if *res != test {
			t.Errorf("currency not matched")
		}
	}
}
