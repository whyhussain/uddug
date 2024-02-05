package main

import (
	"awesomeProject2/model"
	"awesomeProject2/srv"
	"fmt"
	"time"
)

func main() {
	//example
	transaction := []model.Transaction{
		{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
		{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
		{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
		{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
		{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
	}
	//strating the proccess to get last transactions
	tr := srv.Exchange(transaction, "Day")
	fmt.Println(tr)
	for _, v := range tr {
		fmt.Println(v.Value, v.Timestamp.Unix())
	}
}
