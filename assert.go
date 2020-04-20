package assert

import (
	"fmt"
	"testing"
)

type Assert struct {
	t *testing.T
}

func (a Assert) genError(expected, actual interface{}) error {
	return fmt.Errorf("Expected: %+v, Actual: %+v", expected, actual)
}

func (a Assert) EqualsInt(expected, actual int) error {
	if expected == actual {
		return nil
	}
	return a.genError(expected, actual)
}
