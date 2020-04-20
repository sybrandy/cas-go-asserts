package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
)

func TestEqualsInt(t *testing.T) {
	assert := Assert{t: t}

	testCases := []struct {
		expected int
		actual   int
		err      error
	}{
		{1, 1, nil},
		{2, 2, nil},
		{1, 2, fmt.Errorf("Expected: 1, Actual: 2")},
	}

	for _, currCase := range testCases {
		err := assert.EqualsInt(currCase.expected, currCase.actual)
		if !reflect.DeepEqual(err, currCase.err) {
			t.Fatalf("Diff: <%s>", diff.CharacterDiff(currCase.err.Error(), err.Error()))
		}
	}
}
