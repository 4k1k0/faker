package boolean

import "github.com/jaswdr/faker/generator"

// Boolean is a faker struct for Boolean
type Boolean struct {
	generator generator.Generator
}

func NewBoolean(g generator.GeneratorInterface) Boolean {
	return Boolean{generator: generator.NewGenerator(g)}
}

// Bool returns a fake bool for Faker
func (b Boolean) Bool() bool {
	return b.generator.IntBetween(0, 100) > 50
}

// BoolWithChance returns true with a given percentual chance that the value is true, otherwise returns false
func (b Boolean) BoolWithChance(chanceTrue int) bool {
	if chanceTrue <= 0 {
		return false
	} else if chanceTrue >= 100 {
		return true
	}

	return b.generator.IntBetween(0, 100) < chanceTrue
}

// BoolInt returns a fake bool for Integer Boolean
func (b Boolean) BoolInt() int {
	return b.generator.RandomIntElement([]int{0, 1})
}

// BoolString returns a fake bool for string Boolean
func (b Boolean) BoolString(firstArg string, secondArg string) string {
	boolean := []string{firstArg, secondArg}

	return b.generator.RandomStringElement(boolean)
}
