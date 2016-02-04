package gojas

import (
	"testing"
)


func AssertObjectAtPath(t *testing.T, data, path string) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		if ok = jas.IsObjectAt(path); !ok {
			t.Errorf("No json object exists at path[%v]",path)
		}
	} else {
		t.Error("AssertObjectAtPath:Failed to parse test data.")
	}
	return
}

//AssertObjectAtPathWithKeys attempts to locate an object (a JSON doc) at a given path.
// Given that there is an object there, it will assert that ALL the given property keys are found
// in the object, ignoring their value.
func AssertObjectAtPathWithKeys(t *testing.T, data, path string, keys []string) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		if ok = jas.IsObjectAt(path); !ok {
			t.Log("AssertObjectAtPathWithKeys: Asserted path did not find JSON doc.")
			t.Fail()
		} else {
			var obj map[string]interface{}
			obj,_ = jas.objectAtPath(splitPath(path), jas.receptacle)
			for _,k := range keys {
				if _, found := obj[k]; !found {
					t.Log("AssertObjectAtPathWithKeys: Asserted property key was not found in target JSON doc.")
					t.Fail()
					return false
				}
			}
		}
	} else {
		t.Error("AssertObjectAtPathWithKeys: Failed to parse test data.")
	}
	return
}


func AssertBoolAtPath(t *testing.T, data, path string, asserted bool) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		if ok = jas.IsBoolAt(path,asserted); !ok {
			t.Errorf("Bool[%v] not found at path [%v].",asserted,path)
		}
	} else {
		t.Error("AssertBoolAtPath:Failed to parse test data.")
	}
	return
}



func AssertStringAtPath(t *testing.T, data, path, asserted string) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		if ok = jas.IsStringAt(path,asserted); !ok {
			t.Errorf("String[%v] not found at path [%v].",asserted,path)
		}
	} else {
		t.Error("AssertStringAtPath:Failed to parse test data.")
	}
	return
}
func AssertStringArrayAtPath(t *testing.T, data, path string, asserted []interface{}) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		if ok = jas.IsIdenticalStringSliceAt(path,asserted); !ok {
			t.Errorf("String slice [%v] not found at path [%v].",asserted,path)
		}
	} else {
		t.Error("AssertStringArrayAtPath:Failed to parse test data.")
	}
	return
}

func AssertNumberAtPath(t *testing.T, data, path string, asserted float64) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		if ok = jas.IsNumberAt(path,asserted); !ok {
			t.Errorf("Number[%v] not found at path [%v].",asserted,path)
		}
	} else {
		t.Error("AssertNumberAtPath:Failed to parse test data.")
	}
	return
}

func AssertFloatArrayAtPath(t *testing.T, data, path string, asserted []interface{}) (ok bool) {
	if jas, err := MakeJsonAssertion(data); err==nil {
		if ok = jas.IsIdenticalFloatSliceAt(path,asserted); !ok {
			t.Errorf("Number array[%v] not found at path [%v].",asserted,path)
		}
	} else {
		t.Error("AssertFloatArrayAtPath:Failed to parse test data.")
	}
	return
}
