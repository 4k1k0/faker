package tool

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jaswdr/faker/tool"
)

const (
	maxUint = ^uint(0)
	minUint = 0
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

type Generator struct {
	generator tool.GeneratorInterface
	faker     Faker
}

func NewGenerator(g tool.GeneratorInterface) Generator {
	return Generator{
		generator: g,
		faker:     f,
	}
}

// g.RandomDigit returns a fake random digit for Faker
func (g Generator) RandomDigit() int {
	return g.generator.Int() % 10
}

// g.RandomDigitNotNull returns a fake random digit that is not null for Faker
func (g Generator) RandomDigitNotNull() int {
	return g.generator.Int()%8 + 1
}

// g.IntBetween returns a fake Int between a given minimum and maximum values for Faker
func (g Generator) IntBetween(min, max int) int {
	diff := max - min

	if diff < 0 {
		diff = 0
	}

	if diff == 0 {
		return min
	}

	if diff == maxInt {
		return g.generator.Intn(diff)
	}

	return g.generator.Intn(diff+1) + min
}

// g.RandomDigitNot returns a fake random digit for Faker that is not in a list of ignored
func (g Generator) RandomDigitNot(ignore ...int) int {
	inSlice := func(el int, list []int) bool {
		for i := range list {
			if i == el {
				return true
			}
		}

		return false
	}

	for {
		current := g.g.RandomDigit()
		if inSlice(current, ignore) {
			return current
		}
	}
}

// RandomNumber returns a fake random integer number for Faker
func (g Generator) RandomNumber(size int) int {
	if size == 1 {
		return g.RandomDigit()
	}

	var min int = int(math.Pow10(size - 1))
	var max int = int(math.Pow10(size)) - 1

	return g.IntBetween(min, max)
}

// RandomFloat returns a fake random float number for Faker
func (g Generator) RandomFloat(maxDecimals, min, max int) float64 {
	s := fmt.Sprintf("%d.%d", g.IntBetween(min, max-1), g.IntBetween(1, maxDecimals))
	value, _ := strconv.ParseFloat(s, 32)
	return value
}

// Float returns a fake random float number for Faker
func (g Generator) Float(maxDecimals, min, max int) float64 {
	s := fmt.Sprintf("%d.%d", g.IntBetween(min, max-1), g.IntBetween(1, maxDecimals))
	value, _ := strconv.ParseFloat(s, 32)
	return value
}

// Float32 returns a fake random float64 number for Faker
func (g Generator) Float32(maxDecimals, min, max int) float32 {
	s := fmt.Sprintf("%d.%d", g.IntBetween(min, max-1), g.IntBetween(1, maxDecimals))
	value, _ := strconv.ParseFloat(s, 32)
	return float32(value)
}

// Float64 returns a fake random float64 number for Faker
func (g Generator) Float64(maxDecimals, min, max int) float64 {
	s := fmt.Sprintf("%d.%d", g.IntBetween(min, max-1), g.IntBetween(1, maxDecimals))
	value, _ := strconv.ParseFloat(s, 32)
	return float64(value)
}

// Int returns a fake Int number for Faker
func (g Generator) Int() int {
	max := int(^uint(0)>>1) - 1
	min := 0
	return g.IntBetween(min, max)
}

// Int8 returns a fake Int8 number for Faker
func (g Generator) Int8() int8 {
	return int8(g.Int())
}

// Int16 returns a fake Int16 number for Faker
func (g Generator) Int16() int16 {
	return int16(g.Int())
}

// Int32 returns a fake Int32 number for Faker
func (g Generator) Int32() int32 {
	return int32(g.Int())
}

// Int64 returns a fake Int64 number for Faker
func (g Generator) Int64() int64 {
	return int64(g.Int())
}

// UInt returns a fake UInt number for Faker
func (g Generator) UInt() uint {
	maxU := ^uint(0) >> 1
	max := int(maxU)
	return uint(g.IntBetween(0, max))
}

// UInt8 returns a fake UInt8 number for Faker
func (g Generator) UInt8() uint8 {
	return uint8(g.Int())
}

// UInt16 returns a fake UInt16 number for Faker
func (g Generator) UInt16() uint16 {
	return uint16(g.Int())
}

// UInt32 returns a fake UInt32 number for Faker
func (g Generator) UInt32() uint32 {
	return uint32(g.Int())
}

// UInt64 returns a fake UInt64 number for Faker
func (g Generator) UInt64() uint64 {
	return uint64(g.Int())
}

// Int64Between returns a fake Int64 between a given minimum and maximum values for Faker
func (g Generator) Int64Between(min, max int64) int64 {
	return int64(g.IntBetween(int(min), int(max)))
}

// Int32Between returns a fake Int32 between a given minimum and maximum values for Faker
func (g Generator) Int32Between(min, max int32) int32 {
	return int32(g.IntBetween(int(min), int(max)))
}

// Letter returns a fake single letter for Faker
func (g Generator) Letter() string {
	return g.RandomLetter()
}

// RandomLetter returns a fake random string with a random number of letters for Faker
func (g Generator) RandomLetter() string {
	return fmt.Sprintf("%c", g.IntBetween(97, 122))
}

func (g Generator) RandomStringWithLength(l int) string {
	r := []string{}
	for i := 0; i < l; i++ {
		r = append(r, g.RandomLetter())
	}
	return strings.Join(r, "")
}

// RandomStringElement returns a fake random string element from a given list of strings for Faker
func (g Generator) RandomStringElement(s []string) string {
	i := g.IntBetween(0, len(s)-1)
	return s[i]
}

// RandomStringMapKey returns a fake random string key from a given map[string]string for Faker
func (g Generator) RandomStringMapKey(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	i := g.IntBetween(0, len(keys)-1)
	return keys[i]
}

// RandomStringMapValue returns a fake random string value from a given map[string]string for Faker
func (g Generator) RandomStringMapValue(m map[string]string) string {
	values := make([]string, 0, len(m))
	for k := range m {
		values = append(values, m[k])
	}

	i := g.IntBetween(0, len(values)-1)
	return values[i]
}

// RandomIntElement returns a fake random int element form a given list of ints for Faker
func (g Generator) RandomIntElement(a []int) int {
	i := g.IntBetween(0, len(a)-1)
	return a[i]
}

// ShuffleString returns a fake shuffled string from a given string for Faker
func (g Generator) ShuffleString(s string) string {
	orig := strings.Split(s, "")
	dest := make([]string, len(orig))

	for i := 0; i < len(orig); i++ {
		dest[i] = orig[len(orig)-i-1]
	}

	return strings.Join(dest, "")
}

// Numerify returns a fake string that replace all "#" characters with numbers from a given string for Faker
func (g Generator) Numerify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "#" {
			c = strconv.Itoa(g.RandomDigit())
		}

		out = out + c
	}

	return
}

// Lexify  returns a fake string that replace all "?" characters with random letters from a given string for Faker
func (g Generator) Lexify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "?" {
			c = g.RandomLetter()
		}

		out = out + c
	}

	return
}

// Bothify returns a fake string that apply Lexify() and Numerify() on a given string for Faker
func (g Generator) Bothify(in string) (out string) {
	out = g.Lexify(in)
	out = g.Numerify(out)
	return
}

// Asciify   returns a fake string that replace all "*" characters with random ASCII values from a given string for Faker
func (g Generator) Asciify(in string) (out string) {
	for _, c := range strings.Split(in, "") {
		if c == "*" {
			c = fmt.Sprintf("%c", g.IntBetween(97, 126))
		}
		out = out + c
	}

	return
}

// Bool returns a fake bool for Faker
func (g Generator) Bool() bool {
	return g.Boolean().Bool()
}

// BoolWithChance returns true with a given percentual chance that the value is true, otherwise returns false
func (g Generator) BoolWithChance(chanceTrue int) bool {
	return g.Boolean().BoolWithChance(chanceTrue)
}
