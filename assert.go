package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
)

type Assert struct {
	t            *testing.T
	reportErrors bool
}

func (a Assert) logError(expected, actual interface{}) {
	var msg string
	if reflect.TypeOf(expected).String() == "string" {
		a, _ := expected.(string)
		b, _ := actual.(string)
		msg = fmt.Sprintf("Expected: %+v, Actual: %+v, Diff: %s",
			expected, actual, diff.CharacterDiff(a, b))
	} else {
		msg = fmt.Sprintf("Expected: %+v, Actual: %+v", expected, actual)
	}
	if a.reportErrors {
		a.t.Error(msg)
	} else {
		a.t.Log(msg)
	}
}

func (a Assert) Equals(expected, actual interface{}) bool {
	if expected != actual {
		a.logError(expected, actual)
		return false
	}
	return true
}

func (a Assert) EqualsByte(expected, actual byte) bool {
	return a.EqualsUInt8(expected, actual)
}

func (a Assert) EqualsRune(expected, actual rune) bool {
	return a.EqualsInt32(expected, actual)
}

func (a Assert) EqualsInt(expected, actual int) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsInt8(expected, actual int8) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsInt16(expected, actual int16) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsInt32(expected, actual int32) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsInt64(expected, actual int64) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsUInt(expected, actual uint) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsUInt8(expected, actual uint8) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsUInt16(expected, actual uint16) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsUInt32(expected, actual uint32) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsUInt64(expected, actual uint64) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsFloat32(expected, actual float32) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsFloat64(expected, actual float64) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsBool(expected, actual bool) bool {
	return a.Equals(expected, actual)
}

func (a Assert) EqualsString(expected, actual string) bool {
	return a.Equals(expected, actual)
}
