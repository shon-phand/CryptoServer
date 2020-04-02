package services

import (
	"fmt"
	"time"

	"github.com/shon-phand/CryptoServer/sync"
)

func UpdateDatabase() {
	startTime := time.Now()
	curr := sync.SyncCurrency()
	Save(curr)
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
