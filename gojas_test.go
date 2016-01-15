package gojas

import (
	"testing"
)

func TestAssertionMaker(t *testing.T) {
	_, err := MakeJsonAssertion(asserted_json_data)
	if err!=nil {
		t.Fatal("gojas: failed to parse json data")
	}

}

func TestObjectAt(t *testing.T) {
	path := "/user/properties/object/innerObject"

	jas, _ := MakeJsonAssertion(asserted_json_data)

	if !jas.IsObjectAt(path) {
		t.Error("gojas: expected to find 'innerObject' in sample doc")
	}

	bad_path := "/user/properties/object/inner_object"
	if jas.IsObjectAt(bad_path) {
		t.Error("gojas: expected to NOT find 'inner_object' in sample doc")
	}

}

func TestNumberAt(t *testing.T) {
	path := "/user/properties/object/innerObject/baz"
	var val float64 = 11235

	jas, _ := MakeJsonAssertion(asserted_json_data)

	if !jas.IsNumberAt(path, val) {
		t.Error("gojas: expected to find float(11235) in sample doc @ (%v)",path)
	}

	val = val / 2
	if jas.IsNumberAt(path, val) {
		t.Errorf("gojas: expected to NOT find specific float @ (%v)",path)
	}

	bad_path := path+"/notfound"
	if jas.IsNumberAt(bad_path,val) {
		t.Errorf("gojas: expected to NOT find any float @ (%v)",bad_path)
	}

}

func TestBoolAt(t *testing.T) {
	path := "/user/properties/boolean/value"

	jas, _ := MakeJsonAssertion(asserted_json_data)

	if !jas.IsBoolAt(path, true) {
		t.Error("gojas: expected to find bool(true) in sample doc @ (%v)",path)
	}

	if jas.IsBoolAt(path, false){
		t.Errorf("gojas: expected to NOT find bool(false) @ (%v)",path)
	}

	bad_path := path+"/notfound"
	if jas.IsBoolAt(bad_path,true) {
		t.Errorf("gojas: expected to NOT find any bool @ (%v)",bad_path)
	}

}

func TestStringAt(t *testing.T) {
	path := "/user/properties/string/value"
	val := "foobar"

	jas, _ := MakeJsonAssertion(asserted_json_data)

	if !jas.IsStringAt(path, val) {
		t.Error("gojas: expected to find string(foobar) in sample doc @ (%v)",path)
	}

	val = val + "nbat"
	if jas.IsStringAt(path, val) {
		t.Errorf("gojas: expected to NOT find string @ (%v)",path)
	}

	bad_path := path+"/notfound"
	val = "foobar"
	if jas.IsStringAt(bad_path,val) {
		t.Errorf("gojas: expected to NOT find any string @ (%v)",bad_path)
	}

}

func TestStringArrayAt(t *testing.T) {
	path := "/user/properties/stringArray/value"
	original_val := []interface{}{"1", "1", "2", "3", "5", "8"}
	val := original_val

	jas, _ := MakeJsonAssertion(asserted_json_data)

	if !jas.IsIdenticalStringSliceAt(path, val) {
		t.Error("gojas: expected to find []string in sample doc @ (%v)",path)
	}

	val = append(val,"9")
	if jas.IsIdenticalStringSliceAt(path, val) {
		t.Errorf("gojas: expected to NOT find []string @ (%v)",path)
	}

	val = []interface{}{"1","1","1","3","5","8"} // same length, but diff value at index
	if jas.IsIdenticalStringSliceAt(path, val) {
		t.Errorf("Failed to detect unequal slices of strings, of same length @ (%v)",path)
	}

	bad_path := path+"/notfound"
	val = original_val
	if jas.IsIdenticalStringSliceAt(bad_path,val) {
		t.Errorf("gojas: expected to NOT find any []string @ (%v)",bad_path)
	}

}

func TestMatchingStringArrayAt(t *testing.T) {
	path := "/user/properties/stringArray/value"
	original := []interface{}{"1", "1", "2", "3", "5", "8"}
	val := original

	jas, _ := MakeJsonAssertion(asserted_json_data)

	if !jas.IsMatchingStringSliceAt(path, val) {
		t.Error("Error: expected to find []string in sample doc @ (%v)",path)
	}


	val = append(val,"9")
	if jas.IsMatchingStringSliceAt(path, val) {
		t.Errorf("Error: expected to NOT find []string @ (%v)",path)
	}

	val = []interface{}{"1","1","1","3","5","8"} // same length, but diff value at index
	if jas.IsMatchingStringSliceAt(path, val) {
		t.Errorf("Error: Failed to detect unequal slices of strings, of same length @ (%v)",path)
	}


	if !jas.IsMatchingStringSliceAt(path, reverseInterfaceSlice(original)) {
		t.Errorf("Error: Failed to matches slices of strings, of same length but different order @ (%v)",path)
	}

	if !jas.IsMatchingStringSliceAt(path, reverseInterfaceSlice(original)) {
		t.Errorf("Error: Failed to matches slices of strings, of same length but different order @ (%v)", path)
	}

	bad_path := path + "/notfound"
	val = original
	if jas.IsMatchingStringSliceAt(bad_path,val){
		t.Errorf("gojas: expected to NOT find any []string @ (%v)",bad_path)
	}

}

func TestFloatArrayAt(t *testing.T) {
	path := "/user/properties/numberArray/value"
	original_val := []interface{}{1.0, 1.0, 2.0, 3.0, 5.0, 8.0}

	jas, _ := MakeJsonAssertion(asserted_json_data)

	val := original_val
	if !jas.IsIdenticalFloatSliceAt(path, val) {
		t.Error("gojas: expected to find []float64 in sample doc @ (%v)",path)
	}

	val = append(val, 9.0)
	if jas.IsIdenticalFloatSliceAt(path, val) {
		t.Errorf("gojas: expected to NOT find []float64 @ (%v)",path)
	}

	val = []interface{}{1.0, 1.0, 1.0, 3.0, 5.0, 8.0} // same length, diff val at index
	if jas.IsIdenticalFloatSliceAt(path, val) {
		t.Errorf("Failed to detect equal slices of float64s, of same length @ (%v)",path)
	}

	bad_path := path + "/notfound"
	val = original_val
	if jas.IsIdenticalFloatSliceAt(bad_path,val) {
		t.Errorf("gojas: expected to NOT find any []float64 @ (%v)",bad_path)
	}

}

func TestMatchingFloatArrayAt(t *testing.T) {
	path := "/user/properties/numberArray/value"
	original := []interface{}{1.0, 1.0, 2.0, 3.0, 5.0, 8.0}
	val := original

	jas, _ := MakeJsonAssertion(asserted_json_data)

	if !jas.IsMatchingFloatSliceAt(path, val) {
		t.Error("Error: expected to find []float64 in sample doc @path(%v)", path)
	}

	val = append(val, 9.0)
	if jas.IsMatchingFloatSliceAt(path, val) {
		t.Errorf("Error: expected to NOT find []float64 @ (%v)", path)
	}

	val = []interface{}{1.0, 1.0, 1.0, 3.0, 5.0, 8.0} // same length, but diff value at index 2
	if jas.IsMatchingFloatSliceAt(path, val) {
		t.Errorf("Error: Failed to detect unequal slices of float64s, of same length @ (%v)", path)
	}

	if !jas.IsMatchingFloatSliceAt(path, reverseInterfaceSlice(original)) {
		t.Errorf("Error: Failed to matches slices of floats, of same length but different order @ (%v)", path)
	}

	bad_path := path + "/notfound"
	val = original
	if jas.IsMatchingFloatSliceAt(bad_path, val) {
		t.Errorf("gojas: expected to NOT find any []float64 @ (%v)", bad_path)
	}

}
