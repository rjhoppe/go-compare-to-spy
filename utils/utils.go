package utils

import (
	"errors"
	"log"
	"strings"

	"github.com/rjhoppe/go-compare-to-spy/config"
)

// CheckTickerBadChars Tests ticker input for special chars, exported for testing
func CheckTickerBadChars(x string) error {
	intVals := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	specialChars := "!@#$%^&*()-_+={}[]|;:'<>?/~`"
	for _, i := range intVals {
		check := strings.Contains(x, i)
		if check {
			return errors.New("error: invalid ticker - input value contains a number")
		}
	}
	if strings.ContainsAny(x, specialChars) {
		return errors.New("error: invalid ticker - ticker input value contains a symbol")
	}
	return nil
}

// IsTickerValid tests ticker input for valid ticker, exported for testing
func IsTickerValid(ticker string) error {
	_, key, secret, _ := config.Init()
	url := "https://data.alpaca.markets/v2/stocks/" + ticker + "/trades/latest?feed=iex"
	_, err := GetRequest(key, secret, url)
	if err != nil {
		return err
	}
	return nil
}

// TickerValidation checks user ticker input for special chars and valid ticker
func TickerValidation(ticker string) {
	err := CheckTickerBadChars(ticker)
	if err != nil {
		log.Fatal(err)
	}

	err = IsTickerValid(ticker)
	if err != nil {
		log.Fatal(err)
	}
}
