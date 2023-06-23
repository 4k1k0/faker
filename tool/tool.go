package tool

import "testing"

func Expect(t *testing.T, expected, got interface{}, values ...interface{}) {
	t.Helper()
	if expected != got {
		t.Errorf("\nExpected: (%T) %v \nGot:\t  (%T) %v", expected, expected, got, got)
		if len(values) > 0 {
			for _, v := range values {
				t.Errorf("\n%+v", v)
			}
		}

		t.FailNow()
	}
}

func NotExpect(t *testing.T, notExpected, got interface{}, values ...interface{}) {
	t.Helper()
	if notExpected == got {
		t.Errorf("\nNot Expecting: (%T) '%v', but it was", notExpected, notExpected)
		if len(values) > 0 {
			for _, v := range values {
				t.Errorf("\n%+v", v)
			}
		}

		t.FailNow()
	}
}

func ExpectInInt(t *testing.T, expected int, in []int) {
	t.Helper()
	isIn := false
	for _, v := range in {
		if expected == v {
			isIn = true
			break
		}
	}

	if !isIn {
		t.FailNow()
	}
}

func ExpectInString(t *testing.T, expected string, in []string) {
	t.Helper()
	isIn := false
	for _, v := range in {
		if expected == v {
			isIn = true
			break
		}
	}

	if !isIn {
		t.FailNow()
	}
}
