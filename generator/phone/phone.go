package phone

import (
	"fmt"
	"strings"

	"github.com/jaswdr/faker/generator"
	"github.com/jaswdr/faker/tool"
)

// Phone is a faker struct for Phone
type Phone struct {
	generator generator.Generator
}

func NewPhone(g tool.GeneratorInterface) Phone {
	return Phone{generator: generator.NewGenerator(g)}
}

// AreaCode returns a fake area code for Phone
func (p Phone) AreaCode() (code string) {
	number1 := p.generator.IntBetween(2, 9)
	number2 := p.generator.RandomDigit()
	number3 := p.generator.RandomDigitNot(number2)
	return fmt.Sprintf("%d%d%d", number1, number2, number3)
}

// ExchangeCode returns a fake exchange code for Phone
func (p Phone) ExchangeCode() (code string) {
	number1 := p.generator.IntBetween(2, 9)
	number2 := p.generator.RandomDigit()
	number3 := p.generator.RandomDigit()

	if number2 == 1 {
		number3 = p.generator.RandomDigitNot(1)
	}

	return fmt.Sprintf("%d%d%d", number1, number2, number3)
}

// Number returns a fake phone number for Phone
func (p Phone) Number() string {
	number := p.generator.RandomStringElement(phoneFormats)

	// {{areaCode}}
	number = strings.Replace(number, "{{areaCode}}", p.AreaCode(), 1)

	// {{exchangeCode}}
	number = strings.Replace(number, "{{exchangeCode}}", p.ExchangeCode(), 1)

	return p.generator.Numerify(number)
}

// TollFreeAreaCode returns a fake toll free area code for Phone
func (p Phone) TollFreeAreaCode() string {
	return p.generator.RandomStringElement(tollFreeAreaCodes)
}

// ToolFreeNumber returns a fake tool free number for Phone
func (p Phone) ToolFreeNumber() string {
	number := p.generator.RandomStringElement(tollFreeFormats)

	// {{tollFreeAreaCode}}
	number = strings.Replace(number, "{{tollFreeAreaCode}}", p.TollFreeAreaCode(), 1)

	// {{exchangeCode}}
	number = strings.Replace(number, "{{exchangeCode}}", p.ExchangeCode(), 1)

	return p.generator.Numerify(number)
}

// E164Number returns a fake E164 phone number for Phone
func (p Phone) E164Number() string {
	// Excluding 0 since it cannot be any country code in E164
	firstDigit := p.generator.IntBetween(1, 9)
	randomDigits := p.generator.Numerify("##########")

	return fmt.Sprintf("+%d%s", firstDigit, randomDigits)
}
