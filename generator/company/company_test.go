package company

import (
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestCompanyCatchPhrase(t *testing.T) {
	f := faker.New().Company()
	phrase := f.CatchPhrase()
	tool.Expect(t, true, len(phrase) > 0)
	tool.Expect(t, 2, strings.Count(phrase, " ")) // 3 words
}

func TestCompanyBS(t *testing.T) {
	f := faker.New().Company()
	bs := f.BS()
	tool.Expect(t, true, len(bs) > 0)
	tool.Expect(t, 2, strings.Count(bs, " ")) // 3 words
}

func TestCompanySuffix(t *testing.T) {
	f := faker.New().Company()
	value := f.Suffix()
	tool.Expect(t, true, len(value) > 0)
}

func TestCompanyName(t *testing.T) {
	f := faker.New().Company()
	value := f.Name()
	tool.Expect(t, true, len(value) > 0)
}

func TestCompanyJobTitle(t *testing.T) {
	f := faker.New().Company()
	value := f.JobTitle()
	tool.Expect(t, true, len(value) > 0)
}

func TestEIN(t *testing.T) {
	f := faker.New().Company()
	value := f.EIN()
	tool.Expect(t, true, value > 0)
}
