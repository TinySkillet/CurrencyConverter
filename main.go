package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

var (
	currencyFrom string
	currencyTo   string
	amount       string
)

func main() {

	form := huh.NewForm(
		huh.NewGroup(
			// Ask user to select from currency
			huh.NewSelect[string]().
				Title("Choose your base currency").
				Options(
					huh.NewOptions(CurrenciesToShow...)...,
				).Value(&currencyFrom),

			// Ask user to select to currency
			huh.NewSelect[string]().
				Title("Choose your currency to convert to").
				Options(
					huh.NewOptions(CurrenciesToShow...)...,
				).Value(&currencyTo),

			// Ask user for amount
			huh.NewInput().
				Title("Enter amount to convert").
				Value(&amount).
				// amount validation
				Validate(func(str string) error {
					_, err := strconv.ParseFloat(str, 64)
					if err != nil {
						return errors.New("Please enter a float amount!")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	amountInFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Print("Conversion error")
		log.Fatal(err, err.Error())
	}

	var convertedAmt float64

	action := func() {
		Convert(currencyFrom, currencyTo, amountInFloat, &convertedAmt)
	}
	if err := spinner.New().Title("Converting...").Action(action).Run(); err != nil {
		log.Print("What tf is happening")
		log.Fatal(err)
	}
	fmt.Printf("%s %s in %s is %f", amount, currencyFrom, currencyTo, convertedAmt)
}
