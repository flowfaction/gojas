package gojas

import (
//	"log"
	"reflect"
	"strings"
)

//var logme bool
//
//func init() {
//	logme = false
//}

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

func reverseInterfaceSlice(original []interface{}) (reversed []interface{} ){
	reversed = original
	for i, j := 0, len(reversed)-1; i < j; i, j = i+1, j-1 {
        reversed[i], reversed[j] = reversed[j], reversed[i]
    }
	return
}

// compares two slices of interfaces, that are type-asserted to string elements, and compared for length
// and string comparisons, in order.
func areIdenticalStringInterfaceSlices(left []interface{}, right []interface{}) (identical bool) {
	if len(left) == len(right) {

		for i, item := range left {

			if lval, lok := item.(string); lok {

				if rval, rok := right[i].(string); rok {
					if lval != rval {
						return
					}
				} else {
					return
				}

			} else {
				return
			}

		}

		identical = true

	}
	return
}


//Compares two slices of strings, asserted from interface{}. If type assertion fails, or slices are different lengths,
// or values are not the same (order free, using map), it returns false
func areMatchingStringInterfaceSlices(left []interface{}, right []interface{}) (identical bool) {

	if len(left) == len(right) {
		// make the frequency map for the left slice and populate it
		lmap := make(map[string]int,len(left))
		for _, item := range left {
			if value, sok := item.(string); sok {
				lmap[value]=lmap[value]+1
			} else { // not a string, bail
				return
			}
		}

		// make the frequency map for the right slice and populate it
		rmap := make(map[string]int,len(right))
		for _, item := range right {
			if value, sok := item.(string); sok {
				rmap[value]=rmap[value]+1
			} else { // not a string, bail
				return
			}
		}

		// range on left map, using the lkey to look up on rmap
		for lk,lv := range lmap {
			if rv, rok := rmap[lk]; rok {
				// lkey found in rmap, but same frequency?
				if lv!=rv {
					return
				}

			} else { // left key not found in rmap, so bail
				return
			}
		}

		identical = true
	}

	return
}



// compares two slices of interfaces, that are type-asserted to float64 elements, and compared for length
// and float comparisons, in order.
func areIdenticalFloat64InterfaceSlices(left []interface{}, right []interface{}) (identical bool) {
	if len(left) == len(right) {

		for i, item := range left {

			if lval, lok := item.(float64); lok {

				if rval, rok := right[i].(float64); rok {
					if lval != rval {
						return
					}
				} else {
					return
				}

			} else {
				return
			}
		}

		identical = true
	}
	return
}







// recursive
func (jas *JsonAssertion) objectAtPath(path []string, receptacle map[string]interface{}) (sub_map map[string]interface{}, found bool) {

	// local func to clean up the recursion conditional
	// returns true if the key exists and its value is a submap[string]interface{}
	key_and_map := func(key string, m map[string]interface{}) (submap map[string]interface{}, foundkm bool) {
		if sub, ok := m[key]; ok {
//			if logme {
//				log.Printf("type [%v] found at [%v]\n", reflect.TypeOf(sub), key)
//			}
			submap, foundkm = sub.(map[string]interface{})
		} else {
//			if logme {
//				log.Printf("key not found in map (%v)\n", key)
//			}
		}
		return
	}

	if sub_map, found = key_and_map(path[0], receptacle); found {
		if len(path) > 1 {
			return jas.objectAtPath(tail(path), sub_map) // grab the 'tail' of the slice and recurse on the submap
		} // otherwise just return the value of 'found'
	} else {
//		if logme {
//			log.Printf("key (%v) not found, or not of type (map[string]interface{})\n", path[0])
//		}
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
		value, found = something.(float64)
	}
	return
}

//todo extend for other json array element types (that we want to support)
func (jas *JsonAssertion) arrayAtPath(path []string) (value []interface{}, found bool) {
	leaf_map, parent_found := jas.objectAtPath(headUpToLast(path), jas.receptacle)

	isArray := func() bool {
		_, isSlice := leaf_map[last(path)].([]interface{})
//		if isSlice {
//			if logme {
//				log.Printf("type assertion ok:(%v)\n", val)
//			}
//		} else {
//			if logme {
//				log.Printf("FAIL:type assertion failed!\n")
//			}
//		}
		return isSlice
	}

	if parent_found && len(leaf_map) > 0 && isArray() {
		var something interface{}
		something, found = leaf_map[last(path)]
		value, found = something.([]interface{})
	}

//	if logme {
//		log.Printf("was Array? (%v)\n", isArray())
//	}
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
		value, found = something.(bool)
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
		value, found = something.(string)
	}
	return
}
