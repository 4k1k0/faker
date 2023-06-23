package genre

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestGenreName(t *testing.T) {
	v := faker.New().Genre().Name()
	tool.NotExpect(t, "", v)
}

func TestGenreNameWithDescription(t *testing.T) {
	name, description := faker.New().Genre().NameWithDescription()
	tool.NotExpect(t, "", name)
	tool.NotExpect(t, "", description)
}
