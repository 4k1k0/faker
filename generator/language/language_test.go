package language

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestLanguage(t *testing.T) {
	v := faker.New().Language().Language()
	tool.NotExpect(t, "", v)
}

func TestLanguageAbbr(t *testing.T) {
	v := faker.New().Language().LanguageAbbr()
	tool.NotExpect(t, "", v)
}

func TestProgrammingLanguage(t *testing.T) {
	v := faker.New().Language().ProgrammingLanguage()
	tool.NotExpect(t, "", v)
}
