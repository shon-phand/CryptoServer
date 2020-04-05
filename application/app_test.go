package application

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/shon-phand/CryptoServer/controllers"
	"github.com/shon-phand/CryptoServer/domain"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetCurrency(t *testing.T) {
	var test = domain.Currency{"BTC", "6779.37", "6779.03", "6778.12", "6862.59", "6644.30", "6996.27", "USD"}
	var response domain.Currency
	r := gin.Default()
	r.GET("/currency/:symbol", controllers.GetCurrency())
	w := performRequest(r, "GET", "/currency/BTCUSD")
	json.Unmarshal([]byte(w.Body.String()), &response)
	assert.EqualValues(t, http.StatusOK, w.Code)
	if response != test {
		t.Error("error in resposne")
	}

	if w.Code != http.StatusOK {
		t.Error("invalid error code received")
	}
}
