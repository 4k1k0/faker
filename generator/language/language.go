package language

// Language is a faker struct for Language
type Language struct {
}

// Language returns a fake language name for Language
func (l Language) Language() string {
	return l.Faker.RandomStringElement(languages)
}

// LanguageAbbr returns a fake language name for Language
func (l Language) LanguageAbbr() string {
	return l.Faker.RandomStringElement(languagesAbbr)
}

// ProgrammingLanguage returns a fake programming language for Language
func (l Language) ProgrammingLanguage() string {
	return l.Faker.RandomStringElement(programmingLanguagues)
}
