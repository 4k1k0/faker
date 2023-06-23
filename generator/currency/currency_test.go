package currency

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestCurrency(t *testing.T) {
	c := faker.New().Currency()
	currency := c.Currency()
	tool.NotExpect(t, "", currency)
	tool.ExpectInString(t, currency, currencies)
}

func TestCurrencyCode(t *testing.T) {
	c := faker.New().Currency()
	code := c.Code()
	tool.ExpectInString(t, code, currenciesCodes)
}

func TestCurrencyNumber(t *testing.T) {
	c := faker.New().Currency()
	tool.ExpectInInt(t, c.Number(), currenciesNumbers)
}

func TestCurrencyCountry(t *testing.T) {
	c := faker.New().Currency()
	country := c.Country()
	tool.NotExpect(t, "", country)
	tool.ExpectInString(t, country, currenciesCountries)
}

func TestCurrencyAndCode(t *testing.T) {
	c := faker.New().Currency()
	currency, code := c.CurrencyAndCode()
	tool.NotExpect(t, "", currency)
	tool.ExpectInString(t, currency, currencies)
	tool.ExpectInString(t, code, currenciesCodes)
}
