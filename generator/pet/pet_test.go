package pet

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestPetName(t *testing.T) {
	a := faker.New().Pet()
	tool.NotExpect(t, "", a.Name())
}

func TestPetCat(t *testing.T) {
	a := faker.New().Pet()
	tool.NotExpect(t, "", a.Cat())
}

func TestPetDog(t *testing.T) {
	a := faker.New().Pet()
	tool.NotExpect(t, "", a.Dog())
}
