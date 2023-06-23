package pet

// Pet is a faker struct for Pet
type Pet struct {
}

// Dog returns a fake dog name for App
func (p Pet) Dog() string {
	return p.Faker.RandomStringElement(dogNames)
}

// Cat returns a fake cat name for App
func (p Pet) Cat() string {
	return p.Faker.RandomStringElement(catNames)
}

// Name returns a fake pet name for App
func (p Pet) Name() string {
	petNames := []string{}
	petNames = append(petNames, catNames...)
	petNames = append(petNames, dogNames...)
	return p.Faker.RandomStringElement(petNames)
}
