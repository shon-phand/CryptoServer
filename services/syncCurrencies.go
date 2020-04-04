package services

import (
	"fmt"
	"time"

	"github.com/shon-phand/CryptoServer/sync"
	"github.com/shon-phand/CryptoServer/utils/errors"
)

var lock int

func UpdateDatabase() (float64, *errors.RestErr) {

	var TotalTime float64
	if lock == 1 {
		return TotalTime, errors.StatusBadRequestError("synching already in progress")
	}
	lock = 1
	startTime := time.Now()
	curr := sync.SyncCurrency()
	Save(curr)
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
	lock = 0
	return diff.Seconds(), nil
}
