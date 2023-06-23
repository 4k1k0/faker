package address

import (
	"strconv"
	"strings"

	"github.com/jaswdr/faker/generator/person"
	"github.com/jaswdr/faker/tool"
)

// Address is a faker struct for Address
type Address struct {
}

// CityPrefix returns a fake city prefix for Address
func (a Address) CityPrefix() string {
	return tool.RandomStringElement(cityPrefix)
}

// SecondaryAddress returns a fake secondary address for Address
func (a Address) SecondaryAddress() string {
	format := tool.RandomStringElement(secondaryAddressFormats)
	return tool.Bothify(format)
}

// State returns a fake state for Address
func (a Address) State() string {
	return tool.RandomStringElement(state)
}

// StateAbbr returns a fake state abbreviation for Address
func (a Address) StateAbbr() string {
	return tool.RandomStringElement(stateAbbr)
}

// CitySuffix returns a fake city suffix for Address
func (a Address) CitySuffix() string {
	return tool.RandomStringElement(citySuffix)
}

// StreetSuffix returns a fake street suffix for Address
func (a Address) StreetSuffix() string {
	return tool.RandomStringElement(streetSuffix)
}

// BuildingNumber returns a fake building number for Address
func (a Address) BuildingNumber() (bn string) {
	pattern := tool.RandomStringElement(buildingNumber)
	return tool.Numerify(pattern)
}

// City returns a fake city for Address
func (a Address) City() string {
	city := tool.RandomStringElement(cityFormats)

	// {{cityPrefix}}
	city = strings.Replace(city, "{{cityPrefix}}", a.CityPrefix(), 1)

	var p Person = person.Person()

	// {{firstName}}
	city = strings.Replace(city, "{{firstName}}", p.FirstName(), 1)

	// {{lastName}}
	city = strings.Replace(city, "{{lastName}}", p.LastName(), 1)

	// {{citySuffix}}
	city = strings.Replace(city, "{{citySuffix}}", a.CitySuffix(), 1)

	return city
}

// StreetName returns a fake street name for Address
func (a Address) StreetName() string {
	street := tool.RandomStringElement(streetNameFormats)

	var p Person = tool.Person()

	// {{firstName}}
	street = strings.Replace(street, "{{firstName}}", p.FirstName(), 1)

	// {{lastName}}
	street = strings.Replace(street, "{{lastName}}", p.LastName(), 1)

	// {{streetSuffix}}
	street = strings.Replace(street, "{{streetSuffix}}", a.StreetSuffix(), 1)

	return street
}

// StreetAddress returns a fake street address for Address
func (a Address) StreetAddress() string {
	streetAddress := tool.RandomStringElement(streetAddressFormats)

	// {{buildingNumber}}
	streetAddress = strings.Replace(streetAddress, "{{buildingNumber}}", a.BuildingNumber(), 1)

	// {{streetName}}
	streetAddress = strings.Replace(streetAddress, "{{streetName}}", a.StreetName(), 1)

	// {{secondaryAddress}}
	streetAddress = strings.Replace(streetAddress, "{{secondaryAddress}}", a.SecondaryAddress(), 1)

	return streetAddress
}

// PostCode returns a fake postal code for Address
func (a Address) PostCode() string {
	format := tool.RandomStringElement(postCode)
	return tool.Bothify(format)
}

// Address returns a fake Address
func (a Address) Address() string {
	address := tool.RandomStringElement(addressFormats)

	// {{streetAddress}}
	address = strings.Replace(address, "{{streetAddress}}", a.StreetAddress(), 1)

	// {{city}}
	address = strings.Replace(address, "{{city}}", a.City(), 1)

	// {{stateAbbr}}
	address = strings.Replace(address, "{{stateAbbr}}", a.StateAbbr(), 1)

	// {{postCode}}
	address = strings.Replace(address, "{{postCode}}", a.PostCode(), 1)

	return address
}

// Country returns a fake country for Address
func (a Address) Country() string {
	return tool.RandomStringElement(country)
}

// CountryAbbr returns a fake country abbreviation for Address
func (a Address) CountryAbbr() string {
	return tool.RandomStringElement(countryAbbr)
}

// Latitude returns a fake latitude for Address
func (a Address) Latitude() (latitude float64) {
	latitude, _ = strconv.ParseFloat(tool.Numerify("##.######"), 64)
	return
}

// Longitude returns a fake longitude for Address
func (a Address) Longitude() (latitude float64) {
	latitude, _ = strconv.ParseFloat(tool.Numerify("##.######"), 64)
	return
}
