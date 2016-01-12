package gojas

import (
	"encoding/json"
	"strings"
//	"fmt"
)

//JsonAssertion is the struct we use to organize our walking of the JSON doc. The decoder is
// created by the Maker only. At the moment, the assertions are walking the JSON doc each time.
// Consider an extended method set that can reuse a single JsonAssertion, in cases of large numbers of asserts.
type JsonAssertion struct {
	//	Path string
	json       string
	receptacle map[string]interface{}
	decoder    *json.Decoder
}

//MakeJsonAssertion creates and initializes a JsonAssertion, with decoder instance, or returns an error.
func MakeJsonAssertion(data string) (jas *JsonAssertion, err error) {
	jas = &JsonAssertion{json: data, receptacle: make(map[string]interface{})}
	jas.decoder = json.NewDecoder(strings.NewReader(jas.json))
	if err = jas.decoder.Decode(&jas.receptacle); err != nil {
//		if logme{fmt.Printf("ERROR decoding:(%v)\n", err.Error())}
	}
	return
}

func (jas *JsonAssertion) IsObjectAt(path string) (ok bool) {
	_, ok = jas.objectAtPath(splitPath(path), jas.receptacle)
	return
}

func (jas *JsonAssertion) IsNumberAt(path string, val float64) (ok bool) {
	asserted := val
	val, ok = jas.floatAtPath(splitPath(path))
	return ok && val == asserted
}

func (jas *JsonAssertion) IsBoolAt(path string, val bool) (ok bool) {
	asserted := val
	val, ok = jas.boolAtPath(splitPath(path))
	return ok && val == asserted
}

func (jas *JsonAssertion) IsStringAt(path string, val string) (ok bool) {
	asserted := val
	val, ok = jas.stringAtPath(splitPath(path))
	return ok && val == asserted
}

func (jas *JsonAssertion) IsIdenticalFloatSliceAt(path string, val []interface{}) (ok bool) {
	asserted := val
	val, ok = jas.arrayAtPath(splitPath(path))
	return ok && areIdenticalFloat64InterfaceSlices(val,asserted)
}

func (jas *JsonAssertion) IsIdenticalStringSliceAt(path string, val []interface{}) (ok bool) {
	asserted := val
	val, ok = jas.arrayAtPath(splitPath(path))
	return ok && areIdenticalStringInterfaceSlices(val,asserted)
}

// IsMatchingStringArrayAt
// Assert that a string array is found at the given path, which is 'similar' as in:
// same length, same elements but not in the same order.
func (jas *JsonAssertion) IsMatchingStringSliceAt(path string, val []interface{}) (ok bool) {
	asserted := val
	val, ok = jas.arrayAtPath(splitPath(path))
	return ok && areMatchingStringInterfaceSlices(val,asserted)
}



