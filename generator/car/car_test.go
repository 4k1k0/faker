package car

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestCarMaker(t *testing.T) {
	v := faker.New().Car().Maker()
	tool.NotExpect(t, "", v)
}

func TestCarModel(t *testing.T) {
	v := faker.New().Car().Model()
	tool.NotExpect(t, "", v)
}

func TestCarCategory(t *testing.T) {
	v := faker.New().Car().Category()
	tool.NotExpect(t, "", v)
}

func TestCarFuelType(t *testing.T) {
	v := faker.New().Car().FuelType()
	tool.NotExpect(t, "", v)
}

func TestCarTransmissionGear(t *testing.T) {
	v := faker.New().Car().TransmissionGear()
	tool.NotExpect(t, "", v)
}
