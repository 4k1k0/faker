package image

import (
	"fmt"
	"image"
	"io"
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

type ErrorRaiserPngEncoder struct {
	err error
}

func (creator ErrorRaiserPngEncoder) Encode(w io.Writer, m image.Image) error {
	return creator.err
}

func TestImage(t *testing.T) {
	f := faker.New()
	value := f.Image().Image(100, 100)
	tool.Expect(t, fmt.Sprintf("%T", value), "*os.File")
	tool.Expect(t, strings.HasSuffix(value.Name(), ".png"), true, value.Name())
}

func TestImagePanicIfTempFileCreationFails(t *testing.T) {
	f := faker.New()
	img := f.Image()
	expected := fmt.Errorf("temp file creation failed")
	img.TempFileCreator = tool.ErrorRaiserTempFileCreator{err: expected}
	defer func() {
		Expect(t, recover(), expected)
	}()
	img.Image(100, 100)
}

func TestImagePanicIfEncodingFails(t *testing.T) {
	f := faker.New()
	img := f.Image()
	expected := fmt.Errorf("png encoding failed")
	img.PngEncoder = ErrorRaiserPngEncoder{err: expected}
	defer func() {
		tool.Expect(t, recover(), expected)
	}()
	img.Image(100, 100)
}
