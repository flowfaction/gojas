package main

import (
	log "github.com/Sirupsen/logrus"
)



func main() {

	json_data := `{
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

	var ok bool
	var jas *JsonAssertion

	jas, _ = MakeJsonAssertion(json_data)
	path := "/user/properties/object/inner_object"
	ok = jas.AssertObjectAt(path)
	log.Debugf("object found at path [%v], assertion?=%v", path, ok)

	jas, _ = MakeJsonAssertion(json_data)
	path = "/user/properties/object/innerObject/baz"
	var val float64 = 11235
	ok = jas.AssertNumberAt(path, val)
	log.Debugf("number %v found at path [%v], assertion?=%v", val, path, ok)

	jas, _ = MakeJsonAssertion(json_data)
	path = "/user/properties/boolean/value"
	ok = jas.AssertBoolAt(path, true)
	log.Debugf("bool found at path [%v], assertion?=%v", path, ok)

	jas, _ = MakeJsonAssertion(json_data)
	path = "/user/properties/string/value"
	ok = jas.AssertStringAt(path, "foobar")
	log.Debugf("string [%v] found at path [%v], assertion?=%v", "foobar", path, ok)

	jas, _ = MakeJsonAssertion(json_data)
	path = "/user/properties/numberArray/value"
	farr := []interface{}{1.0, 1.0, 2.0, 3.0, 5.0, 8.0} // json numbers are always floats
	ok = jas.AssertFloatArrayAt(path, farr)
	log.Debugf("float array [%v] found at path [%v], assertion?=%v", farr, path, ok)

	jas, _ = MakeJsonAssertion(json_data)
	path = "/user/properties/stringArray/value"
	sarr := []interface{}{"1","1","2","3","5","8"} // json numbers are always floats
	ok = jas.AssertStringArrayAt(path, sarr)
	log.Debugf("string array [%v] found at path [%v], assertion?=%v", sarr, path, ok)


}
