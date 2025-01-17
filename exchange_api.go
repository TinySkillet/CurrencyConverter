package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type CurrencyResponse struct {
	Data map[string]float64 `json:"data"`
}

var CurrenciesToShow = []string{"USD", "EUR", "SGD", "CAD", "INR", "CNY", "JPY"}

func Convert(currencyFrom, currencyTo string, amount float64, convertedAmt *float64) {

	key, err := getApiKey()
	if err != nil {
		log.Fatal(err)
	}
	url := "https://api.freecurrencyapi.com/v1/latest?base_currency=" + currencyFrom + "&currencies=" + strings.Join(CurrenciesToShow, ",")

	client := http.Client{
		Timeout: time.Second * 10,
	}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("apikey", key)

	res, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	var data CurrencyResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	rates := data.Data
	convert_rate := rates[currencyTo]
	*convertedAmt = convert_rate * amount
}

func getApiKey() (string, error) {
	godotenv.Load()
	key := os.Getenv("API_KEY")
	if key == "" {
		return "", errors.New("API_KEY not found in ENVIRONMENT!")
	}
	return key, nil
}
