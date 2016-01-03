package main

import (
	log "github.com/Sirupsen/logrus"
	"encoding/json"
	"strings"
)

func init() {
	log.SetLevel(log.DebugLevel)
}



func SubMapNodeExists(path []string, receptacle map[string]interface{}) (found bool) {

	// returns true if the key exists and its value is a submap[string]interface{}
	key_and_map := func(key string, m map[string]interface{}) (submap map[string]interface{}, foundkm bool) {
		if sub, ok := m[key]; ok {
			submap, foundkm = sub.(map[string]interface{})
		} else {
			log.Debugf("key not found in map (%v)", key)
		}
		return
	}

	var sub_map map[string]interface{}
	if sub_map, found = key_and_map(path[0], receptacle); found {
		if len(path) > 1 {
			return SubMapNodeExists(path[1:], sub_map)
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

	dec := json.NewDecoder(strings.NewReader(json_data))

	var receptacle map[string]interface{}
	if err := dec.Decode(&receptacle); err != nil {
		log.Fatalf("decoding error %v", err.Error())
		//todo: if this error happens, assertion fails, obviously. but don't kill the program
	}

	path := "/user/properties/object/value"

	nodes := strings.Split(path, "/")[1:] // discard the first empty slot due to leading '/'

	// the slice of strings is depth-first navigation of the series of map[string]interface{} parsed by the decoder
	// at each node, represented by a key in the map, the next nested map is checked
	log.Debugf("node found at path[%s]=%v", path, SubMapNodeExists(nodes, receptacle))

}

