package assert

import (
	"fmt"
	"testing"

	"gopkg.in/d4l3k/messagediff.v1"
)

type Assert struct {
	t            *testing.T
	reportErrors bool
}

func (a Assert) logError(expected, actual interface{}, diff string) {
	msg := fmt.Sprintf("Expected: %+v, Actual: %+v, Diff: %s", expected, actual, diff)
	if a.reportErrors {
		a.t.Error(msg)
	} else {
		a.t.Log(msg)
	}
}

func (a Assert) EqualsInt(expected, actual int) bool {
	if diff, equal := messagediff.PrettyDiff(expected, actual); !equal {
		a.logError(expected, actual, diff)
		return false
	}
	return true
}
