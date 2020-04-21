package assert

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
)

type Assert struct {
	t     *testing.T
	level int
}

func (a Assert) logError(expected, actual interface{}) {
	var msg string
	expectedType := "nil"
	actualType := "nil"
	if expected != nil {
		expectedType = reflect.TypeOf(expected).String()
	}
	if actual != nil {
		actualType = reflect.TypeOf(actual).String()
	}
	if expectedType == "string" && actualType == "string" {
		a, _ := expected.(string)
		b, _ := actual.(string)
		msg = fmt.Sprintf("Expected: %+v, Actual: %+v, Diff: %s",
			expected, actual, diff.CharacterDiff(a, b))
	} else {
		msg = fmt.Sprintf("Expected: %+v (%s), Actual: %+v (%s)",
			expected, expectedType, actual, actualType)
	}
	if a.level == 0 {
		a.t.Error(msg)
	} else if a.level == 1 {
		a.t.Log(msg)
	}
}

/*
reflect.Kind Constants
const (
    Invalid Kind = iota
    Bool
    Int
    Int8
    Int16
    Int32
    Int64
    Uint
    Uint8
    Uint16
    Uint32
    Uint64
    Uintptr
    Float32
    Float64
    Complex64
    Complex128
    Array
    Chan
    Func
    Interface
    Map
    Ptr
    Slice
    String
    Struct
    UnsafePointer
)
*/
func (a Assert) isSupported(varKind reflect.Kind) bool {
	return varKind != reflect.Invalid && (varKind <= reflect.Complex128 || varKind == reflect.String)
}

func (a Assert) Equals(expected, actual interface{}) bool {
	if (expected == nil || actual == nil) && expected != actual {
		a.logError(expected, actual)
		return false
	} else if expected == nil && actual == nil {
		return true
	}

	expectedType := reflect.TypeOf(expected)
	actualType := reflect.TypeOf(actual)
	if !a.isSupported(expectedType.Kind()) || !a.isSupported(actualType.Kind()) {
		msg := fmt.Sprintf("Unsupported type in comparison: %s, %s", expectedType.Kind(), actualType.Kind())
		if a.level == 0 {
			a.t.Error(msg)
		} else if a.level == 1 {
			a.t.Log(msg)
		}
		return false
	}

	if expectedType != actualType {
		a.logError(expected, actual)
		return false
	}

	if expected != actual {
		a.logError(expected, actual)
		return false
	}
	return true
}

/*
func (a Assert) EqualsArray(expected, actual, comparitor interface{}) bool {
	a.t.Logf("Comparison function: %s\n", reflect.TypeOf(comparitor))
	exp := reflect.ValueOf(expected)
	act := reflect.ValueOf(actual)

	if exp.Len() != act.Len() {
		msg := fmt.Sprintf("Expected and acutal arrays are of a different length: %d vs. %d", exp.Len(), act.Len())
		if a.reportErrors {
			a.t.Error(msg)
		} else {
			a.t.Log(msg)
		}
		return false
	}

	// think about this more...should be easier...
	switch cmp := comparitor.(type) {
	case func(string, string) bool {
		expVals := expected.([]string)
		actVals := actual.([]string)
		for i, v := range expVals {
			a.EqualsString(v, actVals[i])
		}
	}
	}

	return true
}
*/
