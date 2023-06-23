package currency

import "github.com/jaswdr/faker/tool"

// Currency is a faker struct for generating currency data
type Currency struct {
}

// Currency returns a random currency name
func (c Currency) Currency() string {
	return tool.RandomStringElement(currencies)
}

// Code returns a random currency code
func (c Currency) Code() string {
	return tool.RandomStringElement(currenciesCodes)
}

// Number returns a random currency number
func (c Currency) Number() int {
	return tool.RandomIntElement(currenciesNumbers)
}

// Country returns a random currency country
func (c Currency) Country() string {
	return tool.RandomStringElement(currenciesCountries)
}

// CurrencyAndCode returns a random currency name and currency code
func (c Currency) CurrencyAndCode() (string, string) {
	index := tool.IntBetween(0, len(currencies)-1)
	return currencies[index], currenciesCodes[index]
}
