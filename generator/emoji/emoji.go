package emoji

// Emoji is a faker struct for Emoji
type Emoji struct {
}

// Emoji returns a fake emoji for Emoji
func (a Emoji) Emoji() string {
	return a.Faker.RandomStringElement(emojis)
}

// EmojiCode returns a fake emoji for Emoji
func (a Emoji) EmojiCode() string {
	return a.Faker.RandomStringElement(emojisCode)
}
