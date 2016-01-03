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
	jas.Ok = jas.jsonNumberExists(nodes,jas.receptacle)
	return jas.Ok
}

func (jas *JsonAssertion) AssertObjectAtPath(path string) bool {
	nodes := strings.Split(path, "/")[1:] // discard the first empty slot due to leading '/'
	jas.Ok = jas.jsonObjectExists(nodes,jas.receptacle)
	return jas.Ok
}

func (jas *JsonAssertion) jsonObjectExists(path []string, receptacle map[string]interface{}) (found bool) {

	// local func to clean up the recursion conditional
	// returns true if the key exists and its value is a submap[string]interface{}
	key_and_map := func(key string, m map[string]interface{}) (submap map[string]interface{}, foundkm bool) {
		if sub, ok := m[key]; ok {
			submap, foundkm = sub.(map[string]interface{})
		} else {
			log.Debugf("key not found in map (%v)", key)
		}
		return
	}

	var sub_map map[string]interface{} // declare this var outside of condition, so we can populate the retvar
	if sub_map, found = key_and_map(path[0], receptacle); found {
		if len(path) > 1 {
			return jas.jsonObjectExists(path[1:], sub_map)
		} // otherwise just return the value of 'found'
	} else {
		log.Debugf("key (%v) not found", path[0])
	}

	return
}

func (jas *JsonAssertion) jsonNumberExists(path []string, receptacle map[string]interface{}) (found bool) {

	// local func to clean up the recursion conditional
	// returns true if the key exists and its value is a submap[string]interface{}
	key_and_number := func(key string, m map[string]interface{}) (submap map[string]interface{}, foundkm bool) {
		if sub, ok := m[key]; ok {
			log.Debugf("type at node [%v]",reflect.TypeOf(sub))
			submap, foundkm = sub.(map[string]interface{})
		} else {
			log.Debugf("key not found in map (%v)", key)
		}
		return
	}

	var sub_map map[string]interface{} // declare this var outside of condition, so we can populate the retvar
	if sub_map, found = key_and_number(path[0], receptacle); found {
		if len(path) > 1 {
			return jas.jsonNumberExists(path[1:], sub_map)
		} // otherwise just return the value of 'found'
	} else {
		log.Debugf("key (%v) not found", path[0])
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
						"value": {
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

	jas, _ := MakeJsonAssertion(json_data)

	var ok bool
	path := "/user/properties/object/value"
	ok = jas.AssertObjectAtPath(path)
	log.Debugf("object found at path [%v], ok=%v",path,ok)

	jas, _ = MakeJsonAssertion(json_data)
	path = "/user/properties/object/value/baz"
	ok = jas.AssertNumberAtPath(11235,"/user/properties/object/value/baz")
	log.Debugf("number found at path [%v], ok=%v",path,ok)




}
