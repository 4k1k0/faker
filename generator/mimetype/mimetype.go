package mimetype

// MimeType is a faker struct for MimeType
type MimeType struct {
}

// MimeType returns a fake mime type
func (p MimeType) MimeType() string {
	return p.Faker.RandomStringElement(mimeType)
}
