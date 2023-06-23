package music

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestMusicName(t *testing.T) {
	m := faker.New().Music()
	tool.Expect(t, true, len(m.Name()) > 0)
}

func TestMusicAuthor(t *testing.T) {
	m := faker.New().Music()
	tool.Expect(t, true, len(m.Author()) > 0)
}

func TestMusicGenre(t *testing.T) {
	m := faker.New().Music()
	tool.Expect(t, true, len(m.Genre()) > 0)
}

func TestMusicLength(t *testing.T) {
	m := faker.New().Music()
	tool.Expect(t, true, 2 < m.Length().Minutes() && m.Length().Minutes() < 9)
}
