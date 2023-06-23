package gamer

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestTag(t *testing.T) {
	v := faker.New().Gamer().Tag()
	tool.NotExpect(t, "", v)
}
