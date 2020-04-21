package assert

import (
	"testing"
)

func TestEquals(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected interface{}
		actual   interface{}
		success  bool
	}{
		{int(1), int(1), true},
		{"foo", "foo", true},
		{int(1), "foo", false},
		{int8(1), int(1), false},
		{float32(1), float64(1), false},
		{nil, int(1), false},
		{"foo", nil, false},
		{byte(1), byte(1), true},
		{float32(2), float32(2), true},
		{float64(1), float64(2), false},
		{complex64(2), complex64(2), true},
		{complex128(2), complex128(2), true},
		{complex128(2), complex64(2), false},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "ghi"}, false},
		{"foo", []string{"abc", "ghi"}, false},
		{[]string{"abc", "def", "ghi"}, complex64(123), false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.Equals(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsStringDiff(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected interface{}
		actual   interface{}
		success  bool
	}{
		{"foo", nil, false},
		{"abcd", "abcd", true},
		{"zyx", "zyx", true},
		{"abcd", "zyx", false},
		{"abcd", "abce", false},
		{"abcdefghijklmnop", "abcd efghi jklmnop ", false},
		{"abcdefghijklmnop", "a b c d e f g h i j k l m n o p", false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.Equals(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

/*
func TestEqualsArray(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected   interface{}
		actual     interface{}
		comparitor interface{}
		success    bool
	}{
		{[]string{"abc", "def", "ghi"}, []string{"abc", "def", "ghi"}, assert.EqualsString, true},
		{[]string{"abc", "def", "ghi"}, []string{}, assert.EqualsString, false},
		{[]string{}, []string{"abc", "def", "ghi"}, assert.EqualsString, false},
		{[]string{"abc", "def", "ghi"}, []string{"abc", "ghi"}, assert.EqualsString, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsArray(currCase.expected, currCase.actual, currCase.comparitor) {
			t.Errorf("Assertion failed.")
		}
	}
}
*/
