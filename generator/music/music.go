package music

import (
	"strings"
	"time"

	"github.com/jaswdr/faker"
)

// Music is a faker struct for Music
type Music struct {
}

// Name returns a music name for Music
func (f Music) Name() string {
	var (
		adverb    = f.Faker.RandomStringElement(musicNameAdverbs)
		verb      = f.Faker.RandomStringElement(musicNameVerbs)
		adjective = f.Faker.RandomStringElement(musicNameAdjectives)
		noun      = f.Faker.RandomStringElement(musicNameNouns)
		ending    = f.Faker.RandomStringElement(musicNameEndings)
		name      = musicNameFormats[f.Faker.IntBetween(0, len(musicNameFormats)-1)]
	)

	// {{adverb}}
	name = strings.Replace(name, "{{adverb}}", adverb, 1)

	// {{verb}}
	name = strings.Replace(name, "{{verb}}", verb, 1)

	// {{adjective}}
	name = strings.Replace(name, "{{adjective}}", adjective, 1)

	// {{noun}}
	name = strings.Replace(name, "{{noun}}", noun, 1)

	// {{ending}}
	name = strings.Replace(name, "{{ending}}", ending, 1)

	return name
}

// Author returns the authors name for Music
func (f Music) Author() string {
	p := faker.New().Person()
	return p.Name()
}

// Genre returns genre for Music
func (f Music) Genre() string {
	return f.Faker.RandomStringElement(musicGenres)
}

// Length returns how long the song takes for Music
func (f Music) Length() time.Duration {
	r := f.Faker.IntBetween(128, 512) * int(time.Second)
	return time.Duration(r)
}
