package gender

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestGenderName(t *testing.T) {
	v := faker.New().Gender().Name()
	tool.NotExpect(t, "", v)
	tool.Expect(t, true, v == "masculine" || v == "feminine")
}

func TestGenderAbbr(t *testing.T) {
	v := faker.New().Gender().Abbr()
	tool.NotExpect(t, "", v)
	tool.Expect(t, true, v == "masc" || v == "fem")
}
