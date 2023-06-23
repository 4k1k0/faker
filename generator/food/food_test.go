package food

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestFoodFruit(t *testing.T) {
	v := faker.New().Food().Fruit()
	tool.NotExpect(t, "", v)
}

func TestFoodVegetable(t *testing.T) {
	v := faker.New().Food().Vegetable()
	tool.NotExpect(t, "", v)
}
