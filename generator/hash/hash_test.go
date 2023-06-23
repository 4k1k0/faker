package hash

import (
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jaswdr/faker/tool"
)

// Check if input is hex encoded string
func checkIfHexString(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= 97 && s[i] <= 102) || (s[i] >= 48 && s[i] <= 57)) {
			return false
		}
	}
	return true
}

// tests SHA256()
func TestSHA256(t *testing.T) {
	hash := faker.New().Hash()
	s := hash.SHA256()
	tool.Expect(t, fmt.Sprintf("%T", s), "string")
	tool.Expect(t, 64, len(s))
	tool.Expect(t, true, checkIfHexString(s))
}

// tests SHA512()
func TestSHA512(t *testing.T) {
	hash := faker.New().Hash()
	s := hash.SHA512()
	tool.Expect(t, fmt.Sprintf("%T", s), "string")
	tool.Expect(t, 128, len(s))
	tool.Expect(t, true, checkIfHexString(s))
}

// tests MD5()
func TestMD5(t *testing.T) {
	hash := faker.New().Hash()
	s := hash.MD5()
	tool.Expect(t, fmt.Sprintf("%T", s), "string")
	tool.Expect(t, 32, len(s))
	tool.Expect(t, true, checkIfHexString(s))
}
