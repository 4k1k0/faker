package mimetype

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestMimeType(t *testing.T) {
	p := faker.New().MimeType()
	tool.Expect(t, true, p.MimeType() != "")
}
