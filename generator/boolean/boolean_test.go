package boolean

import (
	"reflect"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

func TestBooleanBool(t *testing.T) {
	f := faker.New().Boolean()
	tp := reflect.TypeOf(f.Bool())
	tool.Expect(t, "bool", tp.String())
}

func TestBooleanBoolWithChance(t *testing.T) {
	f := faker.New().Boolean()
	tp := reflect.TypeOf(f.BoolWithChance(30))
	tool.Expect(t, "bool", tp.String())

	tool.Expect(t, true, f.BoolWithChance(100))
	tool.Expect(t, false, f.BoolWithChance(0))
	tool.Expect(t, true, f.BoolWithChance(101))
	tool.Expect(t, false, f.BoolWithChance(-1))
}

func TestBooleanInt(t *testing.T) {
	p := faker.New().Boolean()
	result := p.BoolInt()
	tool.Expect(t, true, result == 1 || result == 0)
}

func TestBooleanString(t *testing.T) {
	p := faker.New().Boolean()
	args := []string{"yes", "no"}
	result := p.BoolString(args[0], args[1])
	tool.Expect(t, true, result == args[0] || result == args[1])
}
