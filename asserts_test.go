package gojas

import (
	"testing"
//	"fmt"
)

var asserted_json_data string = `{
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


func TestAssertObject(t *testing.T) {

	ok := AssertObjectAtPath(t,asserted_json_data,"/user/properties/object/innerObject")
//	t.Log(t)

	if !ok {
		t.Fatal("Failed to assert existence of known object at known path")
	}

	ok = AssertObjectAtPath(t,asserted_json_data,"/user/properties/object/inner__")
//	t.Log(t)

	if ok {
		t.Fatal("Failed to assert non-existence of object at fake or bad path")
	}

}


func TestAssertNumber(t *testing.T) {

	var val float64 = 11235
	path := "/user/properties/object/innerObject/baz"
	ok := AssertNumberAtPath(t,asserted_json_data,path,val)

	if !ok {
		t.Fatal("Failed to assert matching float value at given path")
	}

	ok = AssertNumberAtPath(t,asserted_json_data,path,val*2)

	if ok {
		t.Fatal("Failed to fail assert of non-matching float value at known path")
	}

}

func TestAssertFloatArray(t *testing.T) {

	original_val := []interface{}{1.0, 1.0, 2.0, 3.0, 5.0, 8.0}

	path := "/user/properties/numberArray/value"
	val := original_val

	ok := AssertFloatArrayAtPath(t,asserted_json_data,path,val)

	if !ok {
		t.Fatal("Failed to pass asserted float array at known path")
	}

	val = append(val,9.0)
	ok = AssertFloatArrayAtPath(t,asserted_json_data,path,val)

	if ok {
		t.Fatal("Failed to fail assert of modified float array at known path")
	}

}

func TestAssertStringArray(t *testing.T) {

	original_val := []interface{}{"1","1","2","3","5","8"}
	val := original_val

	path := "/user/properties/stringArray/value"
	ok := AssertStringArrayAtPath(t,asserted_json_data,path,val)

	if !ok {
		t.Fatal("Failed to pass asserted string array at known path")
	}


	val = append(val,"9")
	ok = AssertStringArrayAtPath(t,asserted_json_data,path,val)

	if ok {
		t.Fatal("Failed to fail assert of modified string array at known path")
	}

}

func TestAssertString(t *testing.T) {
	path := "/user/properties/string/value"

	ok := AssertStringAtPath(t,asserted_json_data,path,"foobar")

	if !ok {
		t.Fatal("Failed to pass assert string('foobar') at known path")
	}

	ok = AssertStringAtPath(t,asserted_json_data,path,"shouldnotbethere")

	if ok {
		t.Fatal("Failed to fail assert of string('shouldnotbethere') at known path")
	}

}

func TestAssertBool(t *testing.T) {
	path := "/user/properties/boolean/value"

	ok := AssertBoolAtPath(t,asserted_json_data,path,true)

	if !ok {
		t.Fatal("Failed to pass assert boolean(true) at known path")
	}

	ok = AssertBoolAtPath(t,asserted_json_data,path,false)

	if ok {
		t.Fatal("Failed to fail assert of not boolean(false) at known path")
	}

}

