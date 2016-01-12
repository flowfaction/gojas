package gojas

import (
	"reflect"
	"strings"
	"fmt"
)

var logme bool

func init() {
	logme=false
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

func tail(slice []string) []string {
	return slice[1:]
}

// compares two slices of interfaces, that are type-asserted to string elements, and compared for length
// and string comparisons, in order.
func isIdenticalStringInterfaceSlices(left []interface{},right []interface{}) (identical bool)  {
	if len(left) == len(right) {
		identical = true
		for i, item := range left {
			lval := item.(string)   // todo: this will panic if fails. need gentler approach for testing
			rval := right[i].(string)
			if lval != rval {
				identical = false
				break
			}
		}
	}
	return identical
}


// compares two slices of interfaces, that are type-asserted to float64 elements, and compared for length
// and float comparisons, in order.
func isIdenticalFloat64InterfaceSlices(left []interface{},right []interface{}) (identical bool)  {
	if len(left) == len(right) {
		identical = true
		for i, item := range left {
			lval := item.(float64)   // todo: this will panic if fails. need gentler approach for testing
			rval := right[i].(float64)
			if lval != rval {
				identical = false
				break
			}
		}
	}
	return identical
}




// recursive
func (jas *JsonAssertion) objectAtPath(path []string, receptacle map[string]interface{}) (sub_map map[string]interface{}, found bool) {

	// local func to clean up the recursion conditional
	// returns true if the key exists and its value is a submap[string]interface{}
	key_and_map := func(key string, m map[string]interface{}) (submap map[string]interface{}, foundkm bool) {
		if sub, ok := m[key]; ok {
			if logme{fmt.Printf("type [%v] found at [%v]\n", reflect.TypeOf(sub), key)}
			submap, foundkm = sub.(map[string]interface{})
		} else {
			if logme{fmt.Printf("key not found in map (%v)\n", key)}
		}
		return
	}

	if sub_map, found = key_and_map(path[0], receptacle); found {
		if len(path) > 1 {
			return jas.objectAtPath(  tail(path) , sub_map) // grab the 'tail' of the slice and recurse on the submap
		} // otherwise just return the value of 'found'
	} else {
		if logme{fmt.Printf("key (%v) not found\n", path[0])}
	}

	return
}

func (jas *JsonAssertion) floatAtPath(path []string) (value float64, found bool) {
	leaf_map, parent_found := jas.objectAtPath(headUpToLast(path), jas.receptacle)
	isFloat := func() bool {
		return reflect.TypeOf(leaf_map[last(path)]).Kind() == reflect.Float64
	}
	if parent_found && len(leaf_map) > 0 && isFloat() {
		var something interface{}
		something, found = leaf_map[last(path)]
		value = something.(float64) // todo: catch all bad type assertion
	}
	return
}

//todo extend for other json array element types (that we want to support)
func (jas *JsonAssertion) arrayAtPath(path []string) (value []interface{}, found bool) {
	leaf_map, parent_found := jas.objectAtPath(headUpToLast(path), jas.receptacle)

	isArray := func() bool {
		val, isSlice := leaf_map[last(path)].([]interface{})
		if isSlice {
			if logme{fmt.Printf("type assertion ok:(%v)\n", val)}
		} else {
			if logme{fmt.Printf("FAIL:type assertion failed!\n")}
		}
		return isSlice
	}

	if parent_found && len(leaf_map) > 0 && isArray() {
		var something interface{}
		something, found = leaf_map[last(path)]
		value = something.([]interface{}) // todo: catch all bad type assertion
	}

	if logme {fmt.Printf("was Array? (%v)\n", isArray())}
	return
}

func (jas *JsonAssertion) boolAtPath(path []string) (value bool, found bool) {
	leaf_map, parent_found := jas.objectAtPath(headUpToLast(path), jas.receptacle)
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

func (jas *JsonAssertion) stringAtPath(path []string) (value string, found bool) {
	leaf_map, parent_found := jas.objectAtPath(headUpToLast(path), jas.receptacle)
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

