package srv

import (
	"awesomeProject2/model"
	"sort"
	"time"
)

var IntervalMap = map[string]time.Duration{
	"Day":  24 * time.Hour,
	"Week": 7 * (24 * time.Hour),
	"Hour": time.Hour,
}

func Exchange(transactions []model.Transaction, interval string) []model.Transaction {

	//first sorting my slice by timestamps
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Timestamp.Before(transactions[j].Timestamp)
	})
	//creating map to save sorted and grouped transactions
	ValuesMap := make(map[int64]int)

	//rounding it by condition and save them into map
	//so cause it was sorted ,map saves only last transactions
	//added 2 conditions cause needed round down by knowiwng month and year
	for _, v := range transactions {
		date := v.Timestamp.Truncate(IntervalMap[interval]).Unix()
		if interval == "Month" {
			date = time.Date(v.Timestamp.Year(), v.Timestamp.Month(), 1, 0, 0, 0, 0, v.Timestamp.Location()).Unix()
		}
		if interval == "Year" {
			date = time.Date(v.Timestamp.Year(), 1, 1, 0, 0, 0, 0, v.Timestamp.Location()).Unix()
		}
		ValuesMap[date] = v.Value
	}

	//returning slice with transactions from map
	transactions = make([]model.Transaction, 0)
	for i, v := range ValuesMap {
		transactions = append(transactions, model.Transaction{
			Value:     v,
			Timestamp: time.Unix(i, 0),
		})
	}
	//sorted 2nd time by timestamp
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Timestamp.Before(transactions[j].Timestamp)
	})

	return transactions
}
