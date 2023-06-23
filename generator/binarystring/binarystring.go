package binarystring

// BinaryString is the faker struct for BinaryString
type BinaryString struct {
}

// BinaryString returns a random binarystring of given input length
func (bn BinaryString) BinaryString(length int) string {
	var bs string
	for i := 0; i < length; i++ {
		bs += bn.faker.RandomStringElement([]string{"0", "1"})
	}
	return bs
}
