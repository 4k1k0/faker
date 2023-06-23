package image

import (
	"image"
	"image/color"
	"os"

	"github.com/jaswdr/faker/tool"
)

// Image is a faker struct for Image
type Image struct {
	TempFileCreator tool.TempFileCreator
	PngEncoder      tool.PngEncoder
}

// Image returns a fake image file
func (i Image) Image(width, height int) *os.File {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	black := color.RGBA{0, 0, 0, 0xff}
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	step := 4
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if y > 0 {
				if x%step == 0 {
					if y%step == 0 {
						img.Set(x, y, black)
					} else {
						img.Set(x, y, white)
					}
				} else {
					img.Set(x, y, white)
				}
			} else {
				img.Set(x, y, white)
			}
		}
	}

	f, err := i.TempFileCreator.TempFile("fake-img-*.png")
	if err != nil {
		panic(err)
	}

	err = i.PngEncoder.Encode(f, img)
	if err != nil {
		panic(err)
	}

	return f
}
