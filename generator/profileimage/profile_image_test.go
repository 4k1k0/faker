package profileimage

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestProfileImage(t *testing.T) {
	f := faker.New()
	value := f.ProfileImage().Image()
	tool.Expect(t, fmt.Sprintf("%T", value), "*os.File")
	tool.Expect(t, strings.HasSuffix(value.Name(), ".jfif"), true, value.Name())
}

func TestProfileImagePanicIfRequestFails(t *testing.T) {
	f := faker.New()
	service := f.ProfileImage()
	expected := fmt.Errorf("request failed")
	service.HTTPClient = tool.ErrorRaiserHTTPClient{err: expected}
	defer func() {
		tool.Expect(t, recover(), expected)
	}()
	service.Image()
}

func TestProfileImagePanicIfTempFileCreationFails(t *testing.T) {
	f := faker.New()
	service := f.ProfileImage()
	expected := fmt.Errorf("temp file creation failed")
	service.TempFileCreator = tool.ErrorRaiserTempFileCreator{err: expected}
	defer func() {
		tool.Expect(t, recover(), expected)
	}()
	service.Image()
}
