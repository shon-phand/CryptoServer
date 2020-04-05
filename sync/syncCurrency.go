package sync

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

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

type Job struct {
	ID     int
	Symbol domain.Symbol
}

type Result struct {
	job      Job
	currency Crnc
}
type Symbols []domain.Symbol
type Tickers domain.Ticker

var (
	symbols Symbols
	jobs    chan Job
	results chan Result
)

func SyncCurrency() Currencies {

	jobs = make(chan Job, 100)
	results = make(chan Result, 100)

	rawSymbols, err := http.Get("https://api.hitbtc.com/api/2/public/symbol")
	if err != nil {
		fmt.Println("error in downloading the symbols")
	}
	defer rawSymbols.Body.Close()

	body, _ := ioutil.ReadAll(rawSymbols.Body)
	err = json.Unmarshal(body, &symbols)
	if err != nil {
		fmt.Println(err.Error())
	}
	go allocate()
	done := make(chan bool)

	go result(done)

	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done

	return AllCurrencies
}

func allocate() {

	for i, v := range symbols {

		job := Job{i, v}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for res := range results {
		//	fmt.Println("appending currency", res.currency.ID, "to all currencies")
		AllCurrencies = append(AllCurrencies, res.currency)
	}
	done <- true
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {

		rawTckr, err := http.Get("https://api.hitbtc.com/api/2/public/ticker/" + job.Symbol.ID)
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
		Response.Symbol = job.Symbol.ID
		Response.ID = job.Symbol.BaseCurrency
		Response.Ask = ticekrs.Ask
		Response.Bid = ticekrs.Bid
		Response.Last = ticekrs.Last
		Response.Open = ticekrs.Open
		Response.Low = ticekrs.Low
		Response.High = ticekrs.High
		Response.FeeCurrency = job.Symbol.FeeCurrency
		output := Result{job, Response}
		results <- output
		//fmt.Println("added currency", Response)
	}
	wg.Done()
}
