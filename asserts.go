// validates unknown bytes as a correct json document - mostly through built-in decoding error generation and checking
// provides higher level validations on the structure of the json documents, but only where amplifying value of built in checking
// executes assertions against the tree through a text dsl, or variadics, resty url equivalents, etc
// takes care of hairy reflection tasks to provide high-value code generation, dsl, _ ...

package gojas

import (
	"testing"
)


// t *testing.T
// func (*T) FailNow
// func (*T) Fail
// func (*T) Fatal
// func (*T) Fatalf
// func (*T) Log
// func (*T) Logf
// func (*T) Error
// func (*T) Errorf


func AssertObjectAtPath(t *testing.T, data, path string) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		ok = jas.IsObjectAt(path)
	} else {
		t.Error("AssertObjectAtPath"+":Failed to parse test data.")
	}
	return
}

func AssertBoolAtPath(t *testing.T, data, path string, asserted bool) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		ok = jas.IsBoolAt(path,asserted)
	} else {
		t.Error("AssertBoolAtPath"+":Failed to parse test data.")
	}
	return
}



func AssertStringAtPath(t *testing.T, data, path, asserted string) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		ok = jas.IsStringAt(path,asserted)
	} else {
		t.Error("AssertStringAtPath"+":Failed to parse test data.")
	}
	return
}
func AssertStringArrayAtPath(t *testing.T, data, path string, asserted []interface{}) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		ok = jas.IsStringArrayAt(path,asserted)
	} else {
		t.Error("AssertStringArrayAtPath"+":Failed to parse test data.")
	}
	return
}

func AssertNumberAtPath(t *testing.T, data, path string, asserted float64) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		ok = jas.IsNumberAt(path,asserted)
	} else {
		t.Error("AssertNumberAtPath"+":Failed to parse test data.")
	}
	return
}

func AssertFloatArrayAtPath(t *testing.T, data, path string, asserted []interface{}) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		ok = jas.IsFloatArrayAt(path,asserted)
	} else {
		t.Error("AssertFloatArrayAtPath"+":Failed to parse test data.")
	}
	return
}
