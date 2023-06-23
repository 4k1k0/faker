package payment

import (
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestCreditCardType(t *testing.T) {
	p := faker.New().Payment()
	tool.Expect(t, true, len(p.CreditCardType()) > 0)
}

func TestCreditCardNumber(t *testing.T) {
	p := faker.New().Payment()
	tool.Expect(t, true, len(p.CreditCardNumber()) > 0)
}

func TestCreditCardExpirationDateString(t *testing.T) {
	p := faker.New().Payment()

	date := p.CreditCardExpirationDateString()
	tool.Expect(t, 5, len(date))
	tool.Expect(t, true, strings.Contains(date, "/"))
}
