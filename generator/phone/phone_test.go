package phone

import (
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestAreaCode(t *testing.T) {
	p := faker.New().Phone()
	tool.Expect(t, true, len(p.AreaCode()) == 3)
}

func TestExchangeCode(t *testing.T) {
	p := faker.New().Phone()
	tool.Expect(t, true, len(p.ExchangeCode()) == 3)
}

func TestNumber(t *testing.T) {
	a := faker.New().Phone()
	number := a.Number()
	tool.Expect(t, true, len(number) > 0)
	tool.Expect(t, false, strings.Contains(number, "{{areaCode}}"))
	tool.Expect(t, false, strings.Contains(number, "{{exchangeCode}}"))
	tool.Expect(t, false, strings.Contains(number, "#"))
	tool.Expect(t, false, strings.Contains(number, "{{"))
	tool.Expect(t, false, strings.Contains(number, "}}"))
}

func TestTollFreeAreaCode(t *testing.T) {
	a := faker.New().Phone()
	code := a.TollFreeAreaCode()
	tool.Expect(t, true, len(code) > 0)
}

func TestTollFreeNumber(t *testing.T) {
	a := faker.New().Phone()
	number := a.ToolFreeNumber()
	tool.Expect(t, true, len(number) > 0)
	tool.Expect(t, false, strings.Contains(number, "{{tollFreeAreaCode}}"))
	tool.Expect(t, false, strings.Contains(number, "{{exchangeCode}}"))
	tool.Expect(t, false, strings.Contains(number, "#"))
	tool.Expect(t, false, strings.Contains(number, "{{"))
	tool.Expect(t, false, strings.Contains(number, "}}"))
}

func TestE164Number(t *testing.T) {
	a := faker.New().Phone()
	number := a.E164Number()
	tool.Expect(t, true, len(number) > 0)
}
