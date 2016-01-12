package gojas

import (
	"testing"
)


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
		ok = jas.IsIdenticalStringArrayAt(path,asserted)
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
