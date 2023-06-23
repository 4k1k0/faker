package faker

import (
	"math/rand"
	"time"

	"github.com/jaswdr/faker/generator/address"
	"github.com/jaswdr/faker/generator/app"
	"github.com/jaswdr/faker/generator/beer"
	"github.com/jaswdr/faker/generator/binarystring"
	"github.com/jaswdr/faker/generator/boolean"
	"github.com/jaswdr/faker/generator/car"
	"github.com/jaswdr/faker/generator/color"
	"github.com/jaswdr/faker/generator/company"
	"github.com/jaswdr/faker/generator/crypto"
	"github.com/jaswdr/faker/generator/currency"
	"github.com/jaswdr/faker/generator/emoji"
	"github.com/jaswdr/faker/generator/file"
	"github.com/jaswdr/faker/generator/food"
	"github.com/jaswdr/faker/generator/gamer"
	"github.com/jaswdr/faker/generator/gender"
	"github.com/jaswdr/faker/generator/genre"
	"github.com/jaswdr/faker/generator/hash"
	"github.com/jaswdr/faker/generator/image"
	"github.com/jaswdr/faker/generator/internet"
	"github.com/jaswdr/faker/generator/language"
	"github.com/jaswdr/faker/generator/lorem"
	"github.com/jaswdr/faker/generator/mimetype"
	"github.com/jaswdr/faker/generator/music"
	"github.com/jaswdr/faker/generator/payment"
	"github.com/jaswdr/faker/generator/person"
	"github.com/jaswdr/faker/generator/pet"
	"github.com/jaswdr/faker/generator/phone"
	"github.com/jaswdr/faker/generator/profileimage"
	"github.com/jaswdr/faker/generator/structgen"
	"github.com/jaswdr/faker/generator/useragent"
	"github.com/jaswdr/faker/generator/uuid"
	"github.com/jaswdr/faker/generator/youtube"
	"github.com/jaswdr/faker/tool"
)

const (
	maxUint = ^uint(0)
	minUint = 0
	maxInt  = int(maxUint >> 1)
	minInt  = -maxInt - 1
)

// Faker is the Generator struct for Faker
type Faker struct {
	Generator tool.GeneratorInterface
}

// Boolean returns a fake Boolean instance for Faker
func (f Faker) Boolean() boolean.Boolean {
	return boolean.NewBoolean(f.Generator)
}

// Map returns a fake Map instance for Faker
func (f Faker) Map() map[string]interface{} {
	m := map[string]interface{}{}
	lorem := f.Lorem()

	randWordType := func() string {
		s := tool.RandomStringElement([]string{"lorem", "bs", "job", "name", "address"})
		switch s {
		case "bs":
			return f.Company().BS()
		case "job":
			return f.Company().JobTitle()
		case "name":
			return f.Person().Name()
		case "address":
			return f.Address().Address()
		}
		return lorem.Word()
	}

	randSlice := func() []string {
		var sl []string
		for ii := 0; ii < f.IntBetween(3, 10); ii++ {
			sl = append(sl, lorem.Word())
		}
		return sl
	}

	for i := 0; i < f.IntBetween(3, 10); i++ {
		t := f.RandomStringElement([]string{"string", "int", "float", "slice", "map"})
		switch t {
		case "string":
			m[lorem.Word()] = randWordType()
		case "int":
			m[lorem.Word()] = f.IntBetween(1, 10000000)
		case "float":
			m[lorem.Word()] = f.Float64(f.IntBetween(1, 4), 1, 1000000)
		case "slice":
			m[lorem.Word()] = randSlice()
		case "map":
			mm := map[string]interface{}{}
			tt := f.RandomStringElement([]string{"string", "int", "float", "slice"})
			switch tt {
			case "string":
				mm[lorem.Word()] = randWordType()
			case "int":
				mm[lorem.Word()] = f.IntBetween(1, 10000000)
			case "float":
				mm[lorem.Word()] = f.Float64(f.IntBetween(1, 4), 1, 1000000)
			case "slice":
				mm[lorem.Word()] = randSlice()
			}
			m[lorem.Word()] = mm
		}
	}

	return m
}

// Lorem returns a fake Lorem instance for Faker
func (f Faker) Lorem() lorem.Lorem {
	return lorem.Lorem{}
}

// Person returns a fake Person instance for Faker
func (f Faker) Person() person.Person {
	return person.Person{}
}

// Address returns a fake Address instance for Faker
func (f Faker) Address() address.Address {
	return address.Address{}
}

// Phone returns a fake Phone instance for Faker
func (f Faker) Phone() phone.Phone {
	return phone.NewPhone(f.generator)
}

// Company returns a fake Company instance for Faker
func (f Faker) Company() company.Company {
	return company.Company{}
}

// Time returns a fake Time instance for Faker
func (f Faker) Time() time.Time {
	return time.Time{}
}

// Internet returns a fake Internet instance for Faker
func (f Faker) Internet() internet.Internet {
	return internet.Internet{}
}

// UserAgent returns a fake UserAgent instance for Faker
func (f Faker) UserAgent() useragent.UserAgent {
	return useragent.UserAgent{}
}

// Payment returns a fake Payment instance for Faker
func (f Faker) Payment() payment.Payment {
	return payment.Payment{}
}

// MimeType returns a fake MimeType instance for Faker
func (f Faker) MimeType() mimetype.MimeType {
	return mimetype.MimeType{}
}

// Color returns a fake Color instance for Faker
func (f Faker) Color() color.Color {
	return color.Color{}
}

// UUID returns a fake UUID instance for Faker
func (f Faker) UUID() uuid.UUID {
	return uuid.UUID{}
}

// Image returns a fake Image instance for Faker
func (f Faker) Image() image.Image {
	return image.Image{tool.TempFileCreatorImpl{}, tool.PngEncoderImpl{}}
}

// File returns a fake File instance for Faker
func (f Faker) File() file.File {
	return file.File{}
}

// YouTube returns a fake YouTube instance for Faker
func (f Faker) YouTube() youtube.YouTube {
	return youtube.YouTube{}
}

// Struct returns a fake Struct instance for Faker
func (f Faker) Struct() structgen.Struct {
	return structgen.Struct{}
}

// Gamer returns a fake Gamer instance for Faker
func (f Faker) Gamer() gamer.Gamer {
	return gamer.Gamer{}
}

// Language returns a fake Language instance for Faker
func (f Faker) Language() language.Language {
	return language.Language{}
}

// Beer returns a fake Beer instance for Faker
func (f Faker) Beer() beer.Beer {
	return beer.Beer{}
}

// Car returns a fake Car instance for Faker
func (f Faker) Car() car.Car {
	return car.Car{}
}

// Food returns a fake Food instance for Faker
func (f Faker) Food() food.Food {
	return food.Food{}
}

// App returns a fake App instance for Faker
func (f Faker) App() app.App {
	return app.App{}
}

// Pet returns a fake Pet instance for Faker
func (f Faker) Pet() pet.Pet {
	return pet.Pet{}
}

// Emoji returns a fake Emoji instance for Faker
func (f Faker) Emoji() emoji.Emoji {
	return emoji.Emoji{}
}

// LoremFlickr returns a fake LoremFlickr instance for Faker
func (f Faker) LoremFlickr() lorem.LoremFlickr {
	return lorem.LoremFlickr{tool.HTTPClientImpl{}, tool.TempFileCreatorImpl{}}
}

// ProfileImage returns a fake ProfileImage instance for Faker
func (f Faker) ProfileImage() profileimage.ProfileImage {
	return profileimage.ProfileImage{tool.HTTPClientImpl{}, tool.TempFileCreatorImpl{}}
}

// Genre returns a fake Genre instance for Faker
func (f Faker) Genre() genre.Genre {
	return genre.Genre{}
}

// Gender returns a fake Gender instance for Faker
func (f Faker) Gender() gender.Gender {
	return gender.Gender{}
}

// BinaryString returns a fake BinaryString instance for Faker
func (f Faker) BinaryString() binarystring.BinaryString {
	return binarystring.BinaryString{}
}

// Hash returns a fake Hash instance for Faker
func (f Faker) Hash() hash.Hash {
	return hash.Hash{}
}

// Music returns a fake Music instance for Faker
func (f Faker) Music() music.Music {
	return music.Music{}
}

// Currency returns a fake Currency instance for Faker
func (f Faker) Currency() currency.Currency {
	return currency.Currency{}
}

// Crypto returns a fake Crypto instance for Faker
func (f Faker) Crypto() crypto.Crypto {
	return crypto.Crypto{}
}

// New returns a new instance of Faker instance with a random seed
func New() (f Faker) {
	seed := rand.NewSource(time.Now().Unix())
	f = NewWithSeed(seed)
	return
}

// NewWithSeed returns a new instance of Faker instance with a given seed
func NewWithSeed(src rand.Source) (f Faker) {
	generator := rand.New(src)
	f = Faker{Generator: generator}
	return
}
