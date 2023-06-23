package beer

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestBeerName(t *testing.T) {
	v := faker.New().Beer().Name()
	tool.NotExpect(t, "", v)
}

func TestBeerStyle(t *testing.T) {
	v := faker.New().Beer().Style()
	tool.NotExpect(t, "", v)
}

func TestBeerHop(t *testing.T) {
	v := faker.New().Beer().Hop()
	tool.NotExpect(t, "", v)
}

func TestBeerMalt(t *testing.T) {
	v := faker.New().Beer().Malt()
	tool.NotExpect(t, "", v)
}

func TestBeerAlcohol(t *testing.T) {
	v := faker.New().Beer().Alcohol()
	tool.NotExpect(t, "", v)
}

func TestBeerIbu(t *testing.T) {
	v := faker.New().Beer().Ibu()
	tool.NotExpect(t, "", v)
}

func TestBeerBlg(t *testing.T) {
	v := faker.New().Beer().Blg()
	tool.NotExpect(t, "", v)
}
