package gamer

// Gamer is a faker struct for Gamer
type Gamer struct {
}

// Tag returns a fake gamer tag for Gamer
func (g Gamer) Tag() string {
	return g.Faker.RandomStringElement(gamerTags)
}
