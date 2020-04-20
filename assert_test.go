package assert

import (
	"testing"
)

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
