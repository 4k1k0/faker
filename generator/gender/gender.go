package gender

// Gender is a faker struct for Gender
type Gender struct {
}

// Name returns a Gender name for Gender
func (f Gender) Name() string {
	return f.Faker.RandomStringElement([]string{"masculine", "feminine"})
}

// Abbr returns a Gender name for Gender
func (f Gender) Abbr() string {
	return f.Faker.RandomStringElement([]string{"masc", "fem"})
}
