package srv

import (
	"awesomeProject2/model"
	"reflect"
	"testing"
	"time"
)

func TestExchangeHour(t *testing.T) {
	transaction := []model.Transaction{
		{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
		{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
		{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
		{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
		{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
	}
	expectedResult := []model.Transaction{
		{Value: 4567, Timestamp: time.Unix(1615870800, 0)},
		{Value: 4321, Timestamp: time.Unix(1615888800, 0)},
		{Value: 5212, Timestamp: time.Unix(1616018400, 0)},
		{Value: 4231, Timestamp: time.Unix(1616022000, 0)},
		{Value: 4456, Timestamp: time.Unix(1616025600, 0)},
	}
	result := Exchange(transaction, "Hour")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Exchange function returned unexpected result.\nExpected: %v\nGot: %v", expectedResult, result)
	}
}
func TestExchangeDay(t *testing.T) {
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

func TestExchangeWeek(t *testing.T) {
	transaction := []model.Transaction{
		{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
		{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
		{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
		{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
		{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
	}
	expectedResult := []model.Transaction{
		{Value: 4456, Timestamp: time.Unix(1615766400, 0)},
	}
	result := Exchange(transaction, "Week")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Exchange function returned unexpected result.\nExpected: %v\nGot: %v", expectedResult, result)
	}

}

func TestExchangeMonth(t *testing.T) {
	transaction := []model.Transaction{
		{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
		{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
		{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
		{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
		{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
	}
	expectedResult := []model.Transaction{
		{Value: 4456, Timestamp: time.Unix(1614553200, 0)},
	}
	result := Exchange(transaction, "Month")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Exchange function returned unexpected result.\nExpected: %v\nGot: %v", expectedResult, result)
	}

}

func TestExchangeYear(t *testing.T) {
	transaction := []model.Transaction{
		{Value: 4456, Timestamp: time.Unix(1616026248, 0)},
		{Value: 4231, Timestamp: time.Unix(1616022648, 0)},
		{Value: 5212, Timestamp: time.Unix(1616019048, 0)},
		{Value: 4321, Timestamp: time.Unix(1615889448, 0)},
		{Value: 4567, Timestamp: time.Unix(1615871448, 0)},
	}
	expectedResult := []model.Transaction{
		{Value: 4456, Timestamp: time.Unix(1609455600, 0)},
	}
	result := Exchange(transaction, "Year")

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Exchange function returned unexpected result.\nExpected: %v\nGot: %v", expectedResult, result)
	}

}
