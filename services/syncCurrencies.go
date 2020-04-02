package services

import "github.com/shon-phand/CryptoServer/sync"

func UpdateDatabase() {
	curr := sync.SyncCurrency()
	Save(curr)

}
