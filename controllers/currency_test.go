package controllers

type CurrencyMock struct{}



TestGetCurrency(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "symbol", Value: "BTC"},
	}

	GetGetCurrency(c)
	assert.EqualValues(t, http.StatusNotFound, response.Code)
}