package lorem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestLoremFlickrImage(t *testing.T) {
	f := faker.New()
	value := f.LoremFlickr().Image(300, 200, []string{}, "", false)
	tool.Expect(t, fmt.Sprintf("%T", value), "*os.File")
	tool.Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}

func TestLoremFlickrImageWithPrefix(t *testing.T) {
	f := faker.New()
	for _, prefix := range []string{"g", "p", "red", "green", "blue"} {
		value := f.LoremFlickr().Image(300, 200, []string{}, prefix, false)
		tool.Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
	}
}

func TestLoremFlickrImageWithCategories(t *testing.T) {
	f := faker.New()
	value := f.LoremFlickr().Image(300, 200, []string{"cat", "dog"}, "", false)
	tool.Expect(t, fmt.Sprintf("%T", value), "*os.File")
	tool.Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())

	value = f.LoremFlickr().Image(300, 200, []string{"cat", "dog"}, "", true)
	tool.Expect(t, fmt.Sprintf("%T", value), "*os.File")
	tool.Expect(t, strings.HasSuffix(value.Name(), ".jpg"), true, value.Name())
}

func TestLoremFlickrImagePanicIfRequestFails(t *testing.T) {
	f := faker.New()
	service := f.LoremFlickr()
	expected := fmt.Errorf("request failed")
	service.HTTPClient = tool.ErrorRaiserHTTPClient{err: expected}
	defer func() {
		tool.Expect(t, recover(), expected)
	}()
	service.Image(300, 200, []string{}, "", false)
}

func TestLoremFlickrImagePanicIfTempFileCreationFails(t *testing.T) {
	f := faker.New()
	service := f.LoremFlickr()
	expected := fmt.Errorf("temp file creation failed")
	service.TempFileCreator = tool.ErrorRaiserTempFileCreator{err: expected}
	defer func() {
		tool.Expect(t, recover(), expected)
	}()
	service.Image(300, 200, []string{}, "", false)
}
