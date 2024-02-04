package srv

import (
	"awesomeProject2/model"
	"sort"
	"time"
)

var IntervalMap = map[string]time.Duration{
	"Day":   24 * time.Hour,
	"Month": 31 * (24 * time.Hour),
	"Week":  7 * (24 * time.Hour),
	"Hour":  time.Hour,
}

func Exchange(transactions []model.Transaction, interval string) []model.Transaction {

	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Timestamp.Before(transactions[j].Timestamp)
	})
	ValuesMap := make(map[int64]int)
	for _, v := range transactions {
		c := v.Timestamp.Truncate(IntervalMap[interval]).Unix()
		ValuesMap[c] = v.Value
	}

	transactions = make([]model.Transaction, 0)
	for i, v := range ValuesMap {
		transactions = append(transactions, model.Transaction{
			Value:     v,
			Timestamp: time.Unix(i, 0),
		})
	}
	return transactions
}
