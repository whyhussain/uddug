
# Тестовое Задание для Uddug
1. Склонировать через __git clone https://github.com/whyhussain/uggug__
2. После в Терминале __cd uddug__
3. Дальше __go mod tidy__
4. Запускаем проект через __go run main.go__

### Немного не понял один момент со структурой данных,может это ошибка

Так в примере у мне структура
#### type Transaction struct {
    Value     int       `json:"value"`
    Timestamp time.Time `json:"time"`
}
#### Но в примере TimeStamp дается в виде int64
### {
    Value : 4456,
    TimeStamp:1616026248,
}

### Поэтому в мейн файле я навсякий показываю его в int64 и в формате __Time__

