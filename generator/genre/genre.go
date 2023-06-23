package genre

// Genre is a faker struct for Genre
type Genre struct {
}

// Name returns a genre name for Genre
func (f Genre) Name() string {
	return f.Faker.RandomStringMapKey(genres)
}

// NameWithDescription returns a name and description for Genre
func (f Genre) NameWithDescription() (string, string) {
	key := f.Faker.RandomStringMapKey(genres)
	return key, genres[key]
}
