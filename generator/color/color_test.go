package color

import (
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestHex(t *testing.T) {
	c := faker.New().Color()

	color := c.Hex()
	tool.Expect(t, 7, len(color))
	tool.Expect(t, true, strings.Contains(color, "#"))
}

func TestRGB(t *testing.T) {
	c := faker.New().Color()

	color := c.RGB()
	tool.Expect(t, true, len(color) >= 5)
	tool.Expect(t, true, len(color) <= 11)
	tool.Expect(t, true, strings.Contains(color, ","))
}

func TestRGBAsArray(t *testing.T) {
	c := faker.New().Color()

	color := c.RGBAsArray()
	tool.Expect(t, 3, len(color))
}

func TestCSS(t *testing.T) {
	c := faker.New().Color()

	color := c.CSS()
	tool.Expect(t, true, len(color) > 10)
	tool.Expect(t, true, len(color) <= 16)
	tool.Expect(t, true, strings.Contains(color, "rgb("))
	tool.Expect(t, true, strings.Contains(color, ","))
	tool.Expect(t, true, strings.Contains(color, ")"))
}

func TestSafeColorName(t *testing.T) {
	c := faker.New().Color()

	color := c.SafeColorName()
	tool.Expect(t, true, len(color) > 0)
}

func TestColorName(t *testing.T) {
	c := faker.New().Color()

	color := c.ColorName()
	tool.Expect(t, true, len(color) > 0)
}
