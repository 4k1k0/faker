package emoji

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestEmoji(t *testing.T) {
	e := faker.New().Emoji()
	tool.NotExpect(t, "", e.Emoji())
}

func TestEmojiCode(t *testing.T) {
	e := faker.New().Emoji()
	tool.NotExpect(t, "", e.EmojiCode())
}
