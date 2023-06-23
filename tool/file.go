package tool

import (
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
)

type ErrorRaiserTempFileCreator struct {
	err error
}

func (creator ErrorRaiserTempFileCreator) TempFile(prefix string) (*os.File, error) {
	return nil, creator.err
}

// TempFileCreator creates temporary files
type TempFileCreator interface {
	TempFile(prefix string) (f *os.File, err error)
}

// TempFileCreatorImpl is the default implementation of TempFileCreator
type TempFileCreatorImpl struct{}

// TempFile creates a temporary file
func (TempFileCreatorImpl) TempFile(prefix string) (f *os.File, err error) {
	return ioutil.TempFile(os.TempDir(), prefix)
}

// PngEncoder encodes a image.Image to a io.Writer
type PngEncoder interface {
	Encode(w io.Writer, m image.Image) error
}

// PngEncoderImpl is the default implementation of the PngEncoder
type PngEncoderImpl struct{}

// Encode does the encoding of the image.Image to an io.Writer
func (encoder PngEncoderImpl) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
