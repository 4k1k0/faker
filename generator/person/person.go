package person

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Person is a faker struct for Person
type Person struct {
}

// ContactInfo struct full of contact info
type ContactInfo struct {
	Phone string
	Email string
}

// Suffix returns a fake suffix for Person
func (p Person) Suffix() string {
	return suffix[p.Faker.IntBetween(0, len(suffix)-1)]
}

// TitleMale returns a fake male title for Person
func (p Person) TitleMale() string {
	return "Mr."
}

// TitleFemale returns a fake female title for Person
func (p Person) TitleFemale() string {
	return "Ms."
}

// GenderMale returns a fake GenderMale for Person
func (p Person) GenderMale() string {
	return "Male"
}

// GenderFemale returns a fake GenderFemale for Person
func (p Person) GenderFemale() string {
	return "Female"
}

// Title returns a fake title for Person
func (p Person) Title() string {
	if p.Faker.IntBetween(0, 1) == 0 {
		return p.TitleMale()
	}

	return p.TitleFemale()
}

// FirstNameMale returns a fake male first mame for Person
func (p Person) FirstNameMale() string {
	index := p.Faker.IntBetween(0, len(firstNameMale)-1)
	return firstNameMale[index]
}

// FirstNameFemale returns a fake female first name for Person
func (p Person) FirstNameFemale() string {
	index := p.Faker.IntBetween(0, len(firstNameFemale)-1)
	return firstNameFemale[index]
}

// FirstName returns a fake first name for Person
func (p Person) FirstName() string {
	names := append(firstNameMale, firstNameFemale...)
	return p.Faker.RandomStringElement(names)
}

// LastName returns a fake last name for Person
func (p Person) LastName() string {
	index := p.Faker.IntBetween(0, len(lastName)-1)
	return lastName[index]
}

// Name returns a fake name for Person
func (p Person) Name() string {
	formats := append(maleNameFormats, femaleNameFormats...)
	name := formats[p.Faker.IntBetween(0, len(formats)-1)]

	// {{titleMale}}
	name = strings.Replace(name, "{{titleMale}}", p.TitleMale(), 1)

	//{{firstNameMale}}
	name = strings.Replace(name, "{{firstNameMale}}", p.FirstNameMale(), 1)

	// {{titleFemale}}
	name = strings.Replace(name, "{{titleFemale}}", p.TitleFemale(), 1)

	//{{firstNameFemale}}
	name = strings.Replace(name, "{{firstNameFemale}}", p.FirstNameFemale(), 1)

	//{{lastName}}
	name = strings.Replace(name, "{{lastName}}", p.LastName(), 1)

	//{{suffix}}
	name = strings.Replace(name, "{{suffix}}", p.Suffix(), 1)

	return name
}

// NameMale returns a fake NameMale for Person
func (p Person) NameMale() string {
	return fmt.Sprintf("%s %s", p.FirstNameMale(), p.LastName())
}

// NameFemale returns a fake NameFemale for Person
func (p Person) NameFemale() string {
	return fmt.Sprintf("%s %s", p.FirstNameFemale(), p.LastName())
}

// Gender returns a fake Gender for Person
func (p Person) Gender() string {
	return p.Faker.RandomStringElement([]string{p.GenderMale(), p.GenderFemale()})
}

// NameAndGender returns a fake NameAndGender for Person
func (p Person) NameAndGender() (string, string) {
	if p.Faker.Boolean().Bool() {
		return p.NameMale(), p.GenderMale()
	}

	return p.NameFemale(), p.GenderFemale()
}

// SSN will generate a random Social Security Number
func (p Person) SSN() string {
	return strconv.Itoa(p.Faker.IntBetween(100000000, 999999999))
}

// Contact will generate a struct with information randomly populated contact information
func (p Person) Contact() ContactInfo {
	return ContactInfo{
		// Tener un constructor de Faker en tools que pueda recibir el tipo de rand
		// y usarle directamente
		// Phone: tools.NewPrivateFaker(p.generator).Phone().Number()
		Phone: p.Faker.Phone().Number(),
		Email: p.Faker.Internet().Email(),
	}
}

// Image return the person profile image
func (p Person) Image() *os.File {
	return p.Faker.ProfileImage().Image()
}
