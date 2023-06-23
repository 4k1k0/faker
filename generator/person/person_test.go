package person

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestTitleMale(t *testing.T) {
	p := faker.New().Person()
	tool.Expect(t, "Mr.", p.TitleMale())
}

func TestTitleFemale(t *testing.T) {
	p := faker.New().Person()
	tool.Expect(t, "Ms.", p.TitleFemale())
}

func TestTitle(t *testing.T) {
	p := faker.New().Person()
	tool.Expect(t, 3, len(p.Title()))
}

func TestSuffix(t *testing.T) {
	p := faker.New().Person()
	suffix := p.Suffix()
	tool.Expect(t, true, len(suffix) > 0)
}

func TestFirstNameMale(t *testing.T) {
	p := faker.New().Person()
	firstName := p.FirstNameMale()
	tool.Expect(t, true, len(firstName) > 0)
}

func TestFirstNameFemale(t *testing.T) {
	p := faker.New().Person()
	firstName := p.FirstNameFemale()
	tool.Expect(t, true, len(firstName) > 0)
}

func TestFirstName(t *testing.T) {
	p := faker.New().Person()
	firstName := p.FirstName()
	tool.Expect(t, true, len(firstName) > 0)
}

func TestLastName(t *testing.T) {
	p := faker.New().Person()
	lastName := p.LastName()
	tool.Expect(t, true, len(lastName) > 0)
}

func TestName(t *testing.T) {
	p := faker.New().Person()
	name := p.Name()
	tool.Expect(t, true, len(name) > 0)
	tool.Expect(t, false, strings.Contains(name, "{{titleMale}}"))
	tool.Expect(t, false, strings.Contains(name, "{{firstNameMale}}"))
	tool.Expect(t, false, strings.Contains(name, "{{titleFemale}}"))
	tool.Expect(t, false, strings.Contains(name, "{{firstNameFemale}}"))
	tool.Expect(t, false, strings.Contains(name, "{{lastName}}"))
	tool.Expect(t, false, strings.Contains(name, "{{suffix}}"))
}

func TestGender(t *testing.T) {
	p := faker.New().Person()
	gender := p.Gender()
	tool.Expect(t, true, gender == "Male" || gender == "Female")
}

func TestGenderMale(t *testing.T) {
	p := faker.New().Person()
	tool.Expect(t, "Male", p.GenderMale())
}

func TestGenderFemale(t *testing.T) {
	p := faker.New().Person()
	tool.Expect(t, "Female", p.GenderFemale())
}

func TestNameAndGender(t *testing.T) {
	p := faker.New().Person()
	name, gender := p.NameAndGender()
	tool.Expect(t, true, name != "")
	tool.Expect(t, true, gender == "Male" || gender == "Female")
}

func TestSSN(t *testing.T) {
	p := faker.New().Person()
	ssn := p.SSN()
	tool.Expect(t, 9, len(ssn))
}

func TestContact(t *testing.T) {
	p := faker.New().Person()
	contact := p.Contact()
	tool.Expect(t, true, len(contact.Phone) > 0)
	tool.Expect(t, true, len(contact.Email) > 0)
}

func TestPersonImage(t *testing.T) {
	p := faker.New().Person()
	image := p.Image()
	tool.Expect(t, fmt.Sprintf("%T", image), "*os.File")
	tool.Expect(t, strings.HasSuffix(image.Name(), ".jfif"), true, image.Name())
}

func TestPersonaNameMale(t *testing.T) {
	p := faker.New().Person()
	name := p.NameMale()
	tool.Expect(t, true, len(name) > 0)
}
