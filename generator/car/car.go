package car

import "github.com/jaswdr/faker/tool"

// Car is a faker struct for Car
type Car struct {
}

// Maker will return a random car maker
func (c Car) Maker() string {
	return tool.RandomStringElement(carMakers)
}

// Model will return a random car model
func (c Car) Model() string {
	return tool.RandomStringElement(carModels)
}

// Category will return a random car category
func (c Car) Category() string {
	return tool.RandomStringElement(carCategories)
}

// FuelType will return a random car fuel type
func (c Car) FuelType() string {
	return tool.RandomStringElement(carFuelTypes)
}

// TransmissionGear will return a random car transmission gear
func (c Car) TransmissionGear() string {
	return tool.RandomStringElement(carTransmissionGears)
}
