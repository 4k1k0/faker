package profileimage

import (
	"io"
	"log"
	"os"

	"github.com/jaswdr/faker/tool"
)

const profileImageBaseURL = "https://thispersondoesnotexist.com/image"

// ProfileImage  is a faker struct for ProfileImage
type ProfileImage struct {
	HTTPClient      tool.HTTPClient
	TempFileCreator tool.TempFileCreator
}

// Image generates a *os.File with a random profile image using the thispersondoesnotexist.com service
func (pi ProfileImage) Image() *os.File {
	resp, err := pi.HTTPClient.Get(profileImageBaseURL)
	if err != nil {
		log.Println("Error while requesting", profileImageBaseURL, ":", err)
		panic(err)
	}

	defer resp.Body.Close()

	f, err := pi.TempFileCreator.TempFile("profil-picture-img-*.jfif")
	if err != nil {
		log.Println("Error while creating a temp file:", err)
		panic(err)
	}

	io.Copy(f, resp.Body)
	return f
}
