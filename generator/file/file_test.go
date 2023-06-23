package file

import (
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestExtension(t *testing.T) {
	p := faker.New().File()
	tool.Expect(t, true, p.Extension() != "")
}

func TestFileWithExtension(t *testing.T) {
	p := faker.New().File()
	tool.Expect(t, true, len(strings.Split(p.FilenameWithExtension(), ".")) == 2)
}
