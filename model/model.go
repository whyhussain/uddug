package model

import "time"

type Transaction struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"time"`
}
