// validates unknown bytes as a correct json document - mostly through built-in decoding error generation and checking
// provides higher level validations on the structure of the json documents, but only where amplifying value of built in checking
// executes assertions against the tree through a text dsl, or variadics, resty url equivalents, etc
// takes care of hairy reflection tasks to provide high-value code generation, dsl, _ ...

package gojas

import (
	"reflect"
	"strings"
	log "github.com/Sirupsen/logrus"
)

//----------------------------------------------------------------------------------------------------------------------
// Internal
//----------------------------------------------------------------------------------------------------------------------

func init() {
	log.SetLevel(log.DebugLevel)
}

func splitPath(path string) []string {
	return strings.Split(path, "/")[1:] // discard the first empty slot due to leading '/'
}

func last(slice []string) string {
	return slice[len(slice)-1]
}

func headUpToLast(slice []string) []string {
	return slice[:len(slice)-1]
}


// recursive
func (jas *JsonAssertion) objectExists(path []string, receptacle map[string]interface{}) (sub_map map[string]interface{}, found bool) {

	// local func to clean up the recursion conditional
	// returns true if the key exists and its value is a submap[string]interface{}
	key_and_map := func(key string, m map[string]interface{}) (submap map[string]interface{}, foundkm bool) {
		if sub, ok := m[key]; ok {
			log.Debugf("type [%v] found at [%v]", reflect.TypeOf(sub), key)
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

func (jas *JsonAssertion) floatExists(path []string) (value float64, found bool) {
	leaf_map, parent_found := jas.objectExists(headUpToLast(path), jas.receptacle)
	isFloat := func() bool {
		return reflect.TypeOf(leaf_map[last(path)]).Kind() == reflect.Float64
	}
	if parent_found && len(leaf_map) > 0 && isFloat() {
		var something interface{}
		something, found = leaf_map[last(path)]
		value = something.(float64)
	}
	return
}

//todo extend for other json array element types (that we want to support)
func (jas *JsonAssertion) arrayExists(path []string) (value []interface{}, found bool) {
	leaf_map, parent_found := jas.objectExists(headUpToLast(path), jas.receptacle)

	isArray := func() bool {
		//		log.Debugf("type of leaf_map item=(%v) @ keyname(%v)",reflect.TypeOf(leaf_map[last(path)]),last(path))
		//		return reflect.TypeOf(leaf_map[last(path)]).Kind()==reflect.SliceOf(reflect.Float64)
		val, isSlice := leaf_map[last(path)].([]interface{})
		if isSlice {
			log.Debugf("type assertion ok:(%v)", val)
		} else {
			log.Debug("FAIL:type assertion failed!")
		}
		return isSlice
	}

	if parent_found && len(leaf_map) > 0 && isArray() {
		var something interface{}
		something, found = leaf_map[last(path)]
		//		log.Debugf("something.type=(%v) /found?(%v)",reflect.TypeOf(something),found)
		value = something.([]interface{})
	}

	log.Debugf("was Array? (%v)", isArray())
	return
}

func (jas *JsonAssertion) boolExists(path []string) (value bool, found bool) {
	leaf_map, parent_found := jas.objectExists(headUpToLast(path), jas.receptacle)
	isBool := func() bool {
		return reflect.TypeOf(leaf_map[last(path)]).Kind() == reflect.Bool
	}
	if parent_found && len(leaf_map) > 0 && isBool() {
		var something interface{}
		something, found = leaf_map[last(path)]
		value = something.(bool)
	}
	return
}

func (jas *JsonAssertion) stringExists(path []string) (value string, found bool) {
	leaf_map, parent_found := jas.objectExists(headUpToLast(path), jas.receptacle)
	isString := func() bool {
		return reflect.TypeOf(leaf_map[last(path)]).Kind() == reflect.String
	}
	if parent_found && len(leaf_map) > 0 && isString() {
		var something interface{}
		something, found = leaf_map[last(path)]
		value = something.(string)
	}
	return
}

