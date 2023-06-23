package youtube

import (
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestGenerateVideoID(t *testing.T) {
	y := faker.New().YouTube()
	tool.Expect(t, 11, len(y.GenerateVideoID()))
}

func TestGenerateFullURL(t *testing.T) {
	y := faker.New().YouTube()
	split := strings.Split(y.GenerateFullURL(), "v=")
	tool.Expect(t, "www.youtube.com/watch?", split[0])
	tool.Expect(t, 11, len(split[1]))
}

func TestGenerateShareURL(t *testing.T) {
	y := faker.New().YouTube()
	split := strings.Split(y.GenerateShareURL(), "/")
	tool.Expect(t, "youtu.be", split[0])
	tool.Expect(t, 11, len(split[1]))
}

func TestGenerateEmbededURL(t *testing.T) {
	y := faker.New().YouTube()
	split := strings.Split(y.GenerateEmbededURL(), "embed/")
	tool.Expect(t, "www.youtube.com/", split[0])
	tool.Expect(t, 11, len(split[1]))
}
