package binarystring

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

// tests BinaryString()
func TestBinaryString(t *testing.T) {
	f := faker.New().BinaryString()
	inputLength := 11
	str := f.BinaryString(inputLength)

	tool.Expect(t, inputLength, len(str))

	isStringValid := true
	for i := 0; i < len(str); i++ {
		if str[i] != '1' && str[i] != '0' {
			isStringValid = false
			break
		}
	}
	tool.Expect(t, true, isStringValid)
}
