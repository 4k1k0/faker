package app

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestAppName(t *testing.T) {
	a := faker.New().App()
	tool.NotExpect(t, "", a.Name())
}

func TestAppVersion(t *testing.T) {
	a := faker.New().App()
	tool.NotExpect(t, "", a.Version())
}
