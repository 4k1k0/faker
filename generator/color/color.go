package color

import (
	"strconv"
	"strings"

	"github.com/jaswdr/faker/tool"
)

// Color is a faker struct for Color
type Color struct {
}

// Hex returns a fake hex for Color
func (c Color) Hex() string {
	color := "#"

	for i := 0; i < 6; i++ {
		color = color + tool.RandomStringElement(colorLetters)
	}

	return color
}

// RGB returns a fake rgb for Color
func (c Color) RGB() string {
	color := strconv.Itoa(tool.IntBetween(0, 255))

	for i := 0; i < 2; i++ {
		color = color + "," + strconv.Itoa(tool.IntBetween(0, 255))
	}

	return color
}

// RGBAsArray returns a fake rgb color in array format for Color
func (c Color) RGBAsArray() [3]string {
	split := strings.Split(c.RGB(), ",")
	return [3]string{split[0], split[1], split[2]}
}

// CSS returns a fake color in CSS format for Color
func (c Color) CSS() string {
	return "rgb(" + c.RGB() + ")"
}

// SafeColorName returns a fake safe color name for Color
func (c Color) SafeColorName() string {
	return tool.RandomStringElement(safeColorNames)
}

// ColorName returns a fake color name for Color
func (c Color) ColorName() string {
	return tool.RandomStringElement(allColorNames)
}
