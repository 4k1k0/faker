package structgen

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

type Basic struct {
	s string
	S string
}

type Nested struct {
	A   string
	B   *Basic
	bar *Basic
}

type BuiltIn struct {
	Uint8   *uint8
	Uint16  *uint16
	Uint32  *uint32
	Uint64  *uint64
	Int     *int
	Int8    *int8
	Int16   *int16
	Int32   *int32
	Int64   *int64
	Float32 *float32
	Float64 *float64
	Bool    *bool
}

type Function struct {
	Number *string `fake:"#"`
	Name   *string `fake:"####"`
	Const  *string `fake:"ABC"`
}

type StructArray struct {
	Bars      []*Basic
	Builds    []BuiltIn
	Skips     []string  `fake:"skip"`
	Strings   []string  `fake:"####" fakesize:"3"`
	SetLen    [5]string `fake:"????"`
	SubStr    [][]string
	SubStrLen [][2]string
	Empty     []*Basic    `fakesize:"0"`
	Multy     []*Function `fakesize:"3"`
}

type NestedArray struct {
	NA []StructArray `fakesize:"2"`
}

func TestStructBasic(t *testing.T) {
	var basic Basic
	faker.New().Struct().Fill(&basic)
	tool.Expect(t, "", basic.s)
	tool.NotExpect(t, "", basic.S)
}

func TestStructNested(t *testing.T) {
	var nested Nested
	faker.New().Struct().Fill(&nested)
	tool.NotExpect(t, "", nested.A)
	tool.NotExpect(t, nil, nested.B)
	tool.NotExpect(t, "", nested.B.S)
	tool.NotExpect(t, nil, nested.bar)
}

func TestStructBuiltInTypes(t *testing.T) {
	var builtIn BuiltIn
	faker.New().Struct().Fill(&builtIn)
	tool.NotExpect(t, nil, builtIn.Uint8)
	tool.NotExpect(t, nil, builtIn.Uint16)
	tool.NotExpect(t, nil, builtIn.Uint32)
	tool.NotExpect(t, nil, builtIn.Uint64)
	tool.NotExpect(t, nil, builtIn.Int)
	tool.NotExpect(t, nil, builtIn.Int8)
	tool.NotExpect(t, nil, builtIn.Int16)
	tool.NotExpect(t, nil, builtIn.Int32)
	tool.NotExpect(t, nil, builtIn.Int64)
	tool.NotExpect(t, nil, builtIn.Float32)
	tool.NotExpect(t, nil, builtIn.Float64)
	tool.NotExpect(t, nil, builtIn.Bool)
}

func TestStructWithFunction(t *testing.T) {
	var function Function
	faker.New().Struct().Fill(&function)
	tool.NotExpect(t, "", function.Number)
	tool.NotExpect(t, "", function.Name)
	tool.NotExpect(t, "ABC", function.Const)
}

func TestStructArray(t *testing.T) {
	var sa StructArray
	faker.New().Struct().Fill(&sa)
	tool.NotExpect(t, 0, sa.Bars)
	tool.NotExpect(t, 0, sa.Builds)
	tool.Expect(t, 3, len(sa.Strings))
	for _, s := range sa.Strings {
		tool.NotExpect(t, "", s)
		tool.Expect(t, 4, len(s))
	}
	tool.Expect(t, 5, len(sa.SetLen))
	for _, s := range sa.SetLen {
		tool.NotExpect(t, "", s)
	}
	for _, s := range sa.SubStr {
		for _, ss := range s {
			tool.NotExpect(t, "", ss)
		}
	}
	for _, s := range sa.SubStrLen {
		tool.Expect(t, 2, len(s))
		for _, ss := range s {
			tool.NotExpect(t, "", ss)
		}
	}
	tool.NotExpect(t, "", sa.Empty)
	tool.NotExpect(t, "", sa.Skips)
	tool.Expect(t, 3, len(sa.Multy))
}

func TestStructNestedArray(t *testing.T) {
	var na NestedArray
	faker.New().Struct().Fill(&na)
	tool.Expect(t, 2, len(na.NA))
	for _, elem := range na.NA {
		tool.NotExpect(t, 0, len(elem.Builds))
		tool.Expect(t, 0, len(elem.Empty))
		tool.Expect(t, 0, len(elem.Skips))
		tool.Expect(t, 3, len(elem.Multy))
	}
}

func TestStructToInt(t *testing.T) {
	var si struct {
		Int         int
		IntConst    int8  `fake:"-123"`
		IntGenerate int64 `fake:"{number:1,10}"`
	}
	faker.New().Struct().Fill(&si)
	tool.NotExpect(t, 0, si.Int)
	tool.NotExpect(t, 0, si.IntConst)
	tool.NotExpect(t, 0, si.IntGenerate)
}

func TestStructToUint(t *testing.T) {
	var su struct {
		Uint         uint
		UintConst    uint8  `fake:"123"`
		UintGenerate uint64 `fake:"{number:1,10}"`
	}
	faker.New().Struct().Fill(&su)
	tool.NotExpect(t, 0, su.Uint)
	tool.NotExpect(t, 0, su.UintConst)
	tool.NotExpect(t, 0, su.UintGenerate)
}

func TestStructToFloat(t *testing.T) {
	var sf struct {
		Float         float32
		FloatConst    float64 `fake:"123.456789"`
		FloatGenerate float32 `fake:"{latitude}"`
	}
	faker.New().Struct().Fill(&sf)
	tool.NotExpect(t, 0, sf.Float)
	tool.NotExpect(t, 0, sf.FloatConst)
	tool.NotExpect(t, 0, sf.FloatGenerate)
}

func TestStructToBool(t *testing.T) {
	var sf struct {
		Bool         bool
		BoolConst    bool `fake:"true"`
		BoolGenerate bool `fake:"{bool}"`
	}
	faker.New().Struct().Fill(&sf)
	tool.NotExpect(t, nil, sf.Bool)
	tool.NotExpect(t, nil, sf.BoolConst)
	tool.NotExpect(t, nil, sf.BoolGenerate)
}
