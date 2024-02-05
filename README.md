# Readme File for Uddug Test Task
1. Clone the repository using the command: `git clone https://github.com/whyhussain/uggug`
2. Navigate to the cloned directory: `cd uddug`
3. Run `go mod tidy` in the terminal.
4. Start the project by running: `go run main.go`

### Issue Understanding Data Structure
I encountered a slight confusion regarding the data structure, and it might be an error.

In the provided example, the structure is defined as follows:
#### type Transaction struct {
    Value     int       `json:"value"`
    Timestamp time.Time `json:"time"`
}
#### However, in the example, `TimeStamp` is given as an int64:
### {
    Value : 4456,
    TimeStamp:1616026248,
}

### Therefore, in the main file, I am explicitly displaying it as int64 and in the format of __Time__.