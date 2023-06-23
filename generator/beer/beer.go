package beer

import (
	"strconv"
)

// Beer is a faker struct for Beer
type Beer struct {
}

// Name will return a random beer name
func (b Beer) Name() string {
	return b.Faker.RandomStringElement(beerNames)
}

// Style will return a random beer style
func (b Beer) Style() string {
	return b.Faker.RandomStringElement(beerStyles)
}

// Hop will return a random beer hop
func (b Beer) Hop() string {
	return b.Faker.RandomStringElement(beerHops)
}

// Malt will return a random beer malt
func (b Beer) Malt() string {
	return b.Faker.RandomStringElement(beerMalts)
}

// Alcohol will return a random beer alcohol level between 2.0 and 10.0
func (b Beer) Alcohol() string {
	return strconv.FormatFloat(b.Faker.RandomFloat(2, 2.0, 10.0), 'f', 1, 64) + "%"
}

// Ibu will return a random beer ibu value between 10 and 100
func (b Beer) Ibu() string {
	return strconv.Itoa(b.Faker.IntBetween(10, 100)) + " IBU"
}

// Blg will return a random beer blg between 5.0 and 20.0
func (b Beer) Blg() string {
	return strconv.FormatFloat(b.Faker.RandomFloat(2, 5.0, 20.0), 'f', 1, 64) + "°Blg"
}
