package food

// Food is a faker struct for Food
type Food struct {
}

// Fruit returns a fake fruit for Food
func (f Food) Fruit() string {
	return f.Faker.RandomStringElement(fruits)
}

// Vegetable returns a fake fruit for Food
func (f Food) Vegetable() string {
	return f.Faker.RandomStringElement(vegetables)
}
