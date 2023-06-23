package file

import (
	"fmt"
)

// File is a faker struct for File
type File struct {
}

// Extension returns a fake Extension file
func (f File) Extension() string {
	return f.Faker.RandomStringElement(extensions)
}

// FilenameWithExtension returns a fake file name with extension
func (f File) FilenameWithExtension() string {
	extension := f.Faker.RandomStringElement(extensions)
	text := f.Faker.Lorem().Word()

	return fmt.Sprintf("%s.%s", text, extension)
}
