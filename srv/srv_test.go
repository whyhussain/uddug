package srv

import (
	"awesomeProject2/model"
	"reflect"
	"testing"
	"time"
)

func TestExchange(t *testing.T) {
	transaction := []model.Transaction{
		{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
		{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
		{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
		{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
		{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
	}
	expectedResult := []model.Transaction{
		{Value: 4321, Timestamp: time.Unix(1615852800, 0)},
		{Value: 4231, Timestamp: time.Unix(1615939200, 0)},
		{Value: 4456, Timestamp: time.Unix(1616025600, 0)},
	}
	result := Exchange(transaction, "Day")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Exchange function returned unexpected result.\nExpected: %v\nGot: %v", expectedResult, result)
	}

}
