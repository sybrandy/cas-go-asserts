package assert

import (
	"fmt"
	"testing"
)

func TestEquals(t *testing.T) {
	assert := Assert{t: t, level: 2}

	testCases := []struct {
		expected interface{}
		actual   interface{}
		success  bool
	}{
		{int(1), int(1), true},
		{nil, nil, true},
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
			t.Errorf("Assertion failed for: %v, %v", currCase.expected, currCase.actual)
		}
	}
}

func TestEqualsStringDiff(t *testing.T) {
	assert := Assert{t: t, level: 2}

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
			t.Errorf("Assertion failed for: %v, %v", currCase.expected, currCase.actual)
		}
	}
}

func TestIsNil(t *testing.T) {
	assert := Assert{t: t, level: 2}

	testCases := []struct {
		actual  interface{}
		success bool
	}{
		{nil, true},
		{"abcd", false},
		{1, false},
		{fmt.Errorf("foo"), false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.IsNil(currCase.actual) {
			t.Errorf("Assertion failed for: %v", currCase.actual)
		}
	}
}

func TestHasError(t *testing.T) {
	assert := Assert{t: t, level: 2}

	testCases := []struct {
		expected string
		actual   error
		success  bool
	}{
		{"", nil, false},
		{"abcd", fmt.Errorf("abcd"), true},
		{"abcd", nil, false},
		{"bar", fmt.Errorf("foo"), false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.HasError(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed for: %v, %v", currCase.expected, currCase.actual)
		}
	}
}

/*
func TestEqualsArray(t *testing.T) {
	assert := Assert{t: t, level: 2}

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
