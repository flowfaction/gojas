package gojas

import (
	"testing"
//	"fmt"
)

var json_data string = `{
		"user":{
			"properties":{
				"string":{
					"type": "string",
					"value": "foobar"
				},
				"number":{
					"type": "number",
					"value": 42
				},
				"boolean":{
					"type": "boolean",
					"value": true
				},
				"numberArray":{
					"type": "array",
					"value": [1,1,2,3,5,8]
				},
				"stringArray":{
					"type": "array",
					"value": ["1","1","2","3","5","8"]
				},
				"object":{
					"type": "object",
					"innerObject": {
						"foo" : "bar",
						"baz": 11235,
						"key" : "value"
					}
				}
			},
			"propertypeer": {
				"peer": 1
			}
		},
		"userpeer" :  {
			"peer" : 1
		}
}`


func TestAssertionMaker(t *testing.T) {
	_, err := MakeJsonAssertion(json_data)
	if err!=nil {
		t.Fatal("gojas: failed to parse json data")
	}

}

func TestObjectAt(t *testing.T) {

	jas, _ := MakeJsonAssertion(json_data)

	path := "/user/properties/object/innerObject"
	ok := jas.IsObjectAt(path)
//	fmt.Printf("object found at path [%v], ok=%v", path, ok)

	if !ok {
		t.Error("gojas: expected to find 'innerObject' in sample doc")
	}

	bad_path := "/user/properties/object/inner_object"
	ok = jas.IsObjectAt(bad_path)
	if ok {
		t.Error("gojas: expected to NOT find 'inner_object' in sample doc")
	}

}


func TestNumberAt(t *testing.T) {

	jas, _ := MakeJsonAssertion(json_data)

	path := "/user/properties/object/innerObject/baz"
	var val float64 = 11235
	ok := jas.IsNumberAt(path, val)

	if !ok {
		t.Error("gojas: expected to find float(11235) in sample doc @ (%v)",path)
	}


	val = val / 2
	ok = jas.IsNumberAt(path, val)
	if ok {
		t.Errorf("gojas: expected to NOT find specific float @ (%v)",path)
	}


	bad_path := path+"/notfound"
	ok = jas.IsNumberAt(bad_path,val)
	if ok {
		t.Errorf("gojas: expected to NOT find any float @ (%v)",bad_path)
	}

}

func TestBoolAt(t *testing.T) {

	jas, _ := MakeJsonAssertion(json_data)

	path := "/user/properties/boolean/value"
	ok := jas.IsBoolAt(path, true)

	if !ok {
		t.Error("gojas: expected to find bool(true) in sample doc @ (%v)",path)
	}


	ok = jas.IsBoolAt(path, false)
	if ok {
		t.Errorf("gojas: expected to NOT find bool(false) @ (%v)",path)
	}


	bad_path := path+"/notfound"
	ok = jas.IsBoolAt(bad_path,true)
	if ok {
		t.Errorf("gojas: expected to NOT find any bool @ (%v)",bad_path)
	}

}


func TestStringAt(t *testing.T) {

	jas, _ := MakeJsonAssertion(json_data)

	path := "/user/properties/string/value"
	val := "foobar"
	ok := jas.IsStringAt(path, val)

	if !ok {
		t.Error("gojas: expected to find string(foobar) in sample doc @ (%v)",path)
	}


	val = val + "nbat"
	ok = jas.IsStringAt(path, val)
	if ok {
		t.Errorf("gojas: expected to NOT find string @ (%v)",path)
	}


	bad_path := path+"/notfound"
	val = "foobar"
	ok = jas.IsStringAt(bad_path,val)
	if ok {
		t.Errorf("gojas: expected to NOT find any string @ (%v)",bad_path)
	}

}

func TestStringArrayAt(t *testing.T) {

	jas, _ := MakeJsonAssertion(json_data)

	original_val := []interface{}{"1","1","2","3","5","8"}

	path := "/user/properties/stringArray/value"
	val := original_val
	ok := jas.IsStringArrayAt(path, val)

	if !ok {
		t.Error("gojas: expected to find []string in sample doc @ (%v)",path)
	}


	val = append(val,"9")
	ok = jas.IsStringArrayAt(path, val)
	if ok {
		t.Errorf("gojas: expected to NOT find []string @ (%v)",path)
	}


	bad_path := path+"/notfound"
	val = original_val
	ok = jas.IsStringArrayAt(bad_path,val)
	if ok {
		t.Errorf("gojas: expected to NOT find any []string @ (%v)",bad_path)
	}

}

func TestFloatArrayAt(t *testing.T) {

	jas, _ := MakeJsonAssertion(json_data)

	original_val := []interface{}{1.0, 1.0, 2.0, 3.0, 5.0, 8.0}

	path := "/user/properties/numberArray/value"
	val := original_val
	ok := jas.IsFloatArrayAt(path, val)

	if !ok {
		t.Error("gojas: expected to find []float64 in sample doc @ (%v)",path)
	}


	val = append(val, 9.0)
	ok = jas.IsFloatArrayAt(path, val)
	if ok {
		t.Errorf("gojas: expected to NOT find []float64 @ (%v)",path)
	}


	bad_path := path+"/notfound"
	val = original_val
	ok = jas.IsFloatArrayAt(bad_path,val)
	if ok {
		t.Errorf("gojas: expected to NOT find any []float64 @ (%v)",bad_path)
	}

}





