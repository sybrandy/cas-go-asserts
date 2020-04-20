package assert

import (
	"testing"
)

func TestEqualsByte(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected byte
		actual   byte
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsByte(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsRune(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected rune
		actual   rune
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsRune(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsInt(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected int
		actual   int
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsInt(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsInt8(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected int8
		actual   int8
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsInt8(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsInt16(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected int16
		actual   int16
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsInt16(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsInt32(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected int32
		actual   int32
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsInt32(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsInt64(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected int64
		actual   int64
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsInt64(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsUInt(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected uint
		actual   uint
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsUInt(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsUInt8(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected uint8
		actual   uint8
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsUInt8(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsUInt16(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected uint16
		actual   uint16
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsUInt16(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsUInt32(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected uint32
		actual   uint32
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsUInt32(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsUInt64(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected uint64
		actual   uint64
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsUInt64(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsFloat32(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected float32
		actual   float32
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsFloat32(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsFloat64(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected float64
		actual   float64
		success  bool
	}{
		{1, 1, true},
		{2, 2, true},
		{1, 2, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsFloat64(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsBool(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected bool
		actual   bool
		success  bool
	}{
		{true, true, true},
		{false, false, true},
		{true, false, false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsBool(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}

func TestEqualsString(t *testing.T) {
	assert := Assert{t: t, reportErrors: false}

	testCases := []struct {
		expected string
		actual   string
		success  bool
	}{
		{"abcd", "abcd", true},
		{"zyx", "zyx", true},
		{"abcd", "zyx", false},
		{"abcd", "abce", false},
	}

	for _, currCase := range testCases {
		if currCase.success != assert.EqualsString(currCase.expected, currCase.actual) {
			t.Errorf("Assertion failed.")
		}
	}
}
