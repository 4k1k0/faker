package lorem

import (
	"io"
	"log"
	"os"
	"strconv"

	"github.com/jaswdr/faker/tool"
)

// LoremFlickr is a faker struct for LoremFlickr
type LoremFlickr struct {
	HTTPClient      tool.HTTPClient
	TempFileCreator tool.TempFileCreator
}

// Image generates a *os.File with a random image using the loremflickr.com service
func (lf LoremFlickr) Image(width, height int, categories []string, prefix string, categoriesStrict bool) *os.File {

	url := loremFlickrBaseURL

	switch prefix {
	case "g":
		url += "/g"
	case "p":
		url += "/p"
	case "red":
		url += "/red"
	case "green":
		url += "/green"
	case "blue":
		url += "/blue"
	}

	url += string('/') + strconv.Itoa(width) + string('/') + strconv.Itoa(height)

	if len(categories) > 0 {

		url += string('/')

		for _, category := range categories {
			url += category + string(',')
		}

		if categoriesStrict {
			url += "/all"
		}
	}

	resp, err := lf.HTTPClient.Get(url)
	if err != nil {
		log.Println("Error while requesting", url, ":", err)
		panic(err)
	}

	defer resp.Body.Close()
	f, err := lf.TempFileCreator.TempFile("loremflickr-img-*.jpg")
	if err != nil {
		log.Println("Error while creating a temp file:", err)
		panic(err)
	}

	io.Copy(f, resp.Body)
	return f
}
