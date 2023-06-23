package company

import (
	"strings"
)

// Company is a faker struct for Company
type Company struct {
}

// CatchPhrase returns a fake catch phrase for Company
func (c Company) CatchPhrase() (phrase string) {
	for i, words := range catchPhraseWords {
		if i > 0 {
			phrase += " "
		}
		phrase += tool.RandomStringElement(words)
	}

	return
}

// BS returns a fake bs words for Company
func (c Company) BS() (bs string) {
	for i, words := range bsWords {
		if i > 0 {
			bs += " "
		}
		bs += tool.RandomStringElement(words)
	}

	return
}

// Suffix returns a fake suffix for Company
func (c Company) Suffix() string {
	return tool.RandomStringElement(companySuffix)
}

// Name returns a fake name for Company
func (c Company) Name() string {
	name := tool.RandomStringElement(companyNameFormat)

	// {{companySuffix}}
	name = strings.Replace(name, "{{companySuffix}}", c.Suffix(), 1)

	// {{lastName}}
	p := tool.Person()
	name = strings.Replace(name, "{{lastName}}", p.LastName(), 3)

	return name
}

// JobTitle returns a fake job title for Company
func (c Company) JobTitle() string {
	return tool.RandomStringElement(jobTitle)
}

// EIN returns a fake EIN codes for Company
func (c Company) EIN() int {
	return tool.RandomIntElement(einPrefixes)
}
