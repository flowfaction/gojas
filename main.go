package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"strings"
	"reflect"
)

func init() {
	log.SetLevel(log.DebugLevel)
}



type JsonAssertion struct {
	//	Path string
	Ok         bool
	json       string
	receptacle map[string]interface{}
	decoder    *json.Decoder
}

func MakeJsonAssertion(data string) (jas *JsonAssertion, err error) {
	jas = &JsonAssertion{json: data,receptacle:make(map[string]interface{})}
	jas.decoder = json.NewDecoder(strings.NewReader(jas.json))
	if err = jas.decoder.Decode(&jas.receptacle); err != nil {
		log.Errorf("decoding error %v", err.Error())
	}
	return
}

func (jas *JsonAssertion) AssertNumberAtPath(number int, path string) bool {
	nodes := strings.Split(path, "/")[1:] // discard the first empty slot due to leading '/'
	_,jas.Ok = jas.floatExists(nodes,jas.receptacle)
	return jas.Ok
}

func (jas *JsonAssertion) AssertObjectAtPath(path string) bool {
	nodes := strings.Split(path, "/")[1:] // discard the first empty slot due to leading '/'
	_ , jas.Ok = jas.objectExists(nodes,jas.receptacle)
	return jas.Ok
}

func (jas *JsonAssertion) AssertBoolAtPath(tf bool,path string) bool {
	nodes := strings.Split(path, "/")[1:] // discard the first empty slot due to leading '/'
	_ , jas.Ok = jas.boolExists(nodes,jas.receptacle)
	return jas.Ok
}




func last(slice []string) string {
	return slice[len(slice)-1]
}

func headUpToLast(slice []string) []string {
	return slice[:len(slice)-1]
}


func (jas *JsonAssertion) objectExists(path []string, receptacle map[string]interface{}) (sub_map map[string]interface{}, found bool) {

	// local func to clean up the recursion conditional
	// returns true if the key exists and its value is a submap[string]interface{}
	key_and_map := func(key string, m map[string]interface{}) (submap map[string]interface{}, foundkm bool) {
		if sub, ok := m[key]; ok {
			log.Debugf("type [%v] found at [%v]",reflect.TypeOf(sub),key)
			submap, foundkm = sub.(map[string]interface{})
		} else {
			log.Debugf("key not found in map (%v)", key)
		}
		return
	}

	if sub_map, found = key_and_map(path[0], receptacle); found {
		if len(path) > 1 {
			return jas.objectExists(path[1:], sub_map)
		} // otherwise just return the value of 'found'
	} else {
		log.Debugf("key (%v) not found", path[0])
	}

	return
}


func (jas *JsonAssertion) floatExists(path []string, receptacle map[string]interface{}) (value float64,found bool) {
	leaf_map, parent_found := jas.objectExists(headUpToLast(path),jas.receptacle)
	isFloat := func() bool {
		return reflect.TypeOf(leaf_map[last(path)]).Kind()==reflect.Float64
	}
	if parent_found && len(leaf_map)>0 && isFloat() {
		var something interface{}
		something, found = leaf_map[last(path)]
		value = something.(float64)
	}
	return
}

func (jas *JsonAssertion) boolExists(path []string, receptacle map[string]interface{}) (value bool,found bool) {
	leaf_map, parent_found := jas.objectExists(headUpToLast(path),jas.receptacle)
	isBool := func() bool {
		return reflect.TypeOf(leaf_map[last(path)]).Kind()==reflect.Bool
	}
	if parent_found && len(leaf_map)>0 && isBool() {
		var something interface{}
		something, found = leaf_map[last(path)]
		value = something.(bool)
	}
	return
}



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
					"array":{
						"type": "array",
						"value": [1,1,2,3,5,8]
					},
					"object":{
						"type": "object",
						"inner_object": {
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
	path := "/user/properties/object/value"
	ok = jas.AssertObjectAtPath(path)
	log.Debugf("object found at path [%v], ok=%v",path,ok)

	jas, _ = MakeJsonAssertion(json_data)
	path = "/user/properties/object/inner_object/baz"
	ok = jas.AssertNumberAtPath(11235,path)
	log.Debugf("number found at path [%v], ok=%v",path,ok)


	jas, _ = MakeJsonAssertion(json_data)
	path = "/user/properties/boolean/value"
	ok = jas.AssertBoolAtPath(true,path)
	log.Debugf("bool found at path [%v], ok=%v",path,ok)



}
