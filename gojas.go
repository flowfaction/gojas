package gojas

import (
	"encoding/json"
	"strings"
	"fmt"
)

//----------------------------------------------------------------------------------------------------------------------
// Exported
//----------------------------------------------------------------------------------------------------------------------


type JsonAssertion struct {
	//	Path string
	json       string
	receptacle map[string]interface{}
	decoder    *json.Decoder
}

func MakeJsonAssertion(data string) (jas *JsonAssertion, err error) {
	jas = &JsonAssertion{json: data, receptacle: make(map[string]interface{})}
	jas.decoder = json.NewDecoder(strings.NewReader(jas.json))
	if err = jas.decoder.Decode(&jas.receptacle); err != nil {
		fmt.Printf("ERROR decoding:(%v)", err.Error())
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

func (jas *JsonAssertion) IsFloatArrayAt(path string, val []interface{}) (ok bool) {
	asserted := val
	val, ok = jas.arrayAtPath(splitPath(path))

	arraysAreIdentical := func() (matches bool) {
		if len(val) == len(asserted) {
			matches = true
			for i, item := range val {
				aval := item.(float64)
				bval := asserted[i].(float64)
				if aval != bval {
					matches = false
					break
				}
			}
		}
		return
	}

	return ok && arraysAreIdentical()
}

func (jas *JsonAssertion) IsIdenticalStringArrayAt(path string, val []interface{}) (ok bool) {
	asserted := val
	val, ok = jas.arrayAtPath(splitPath(path))

	arraysMatch := func() (matches bool) {
		if len(val) == len(asserted) {
			matches = true
			for i, item := range val {
				aval := item.(string)   // todo: this will panic if fails. need gentler approach for testings
				bval := asserted[i].(string)
				if strings.Compare(aval, bval)!=0 {
					matches = false
					break
				}
			}
		}
		return
	}
	// val comes from json retrieval, asserted comes from 'client' code
	return ok && arraysMatch()
}

// IsMatchingStringArrayAt
// Assert that a string array is found at the given path, which is 'similar' as in:
// same length, same elements but not in the same order.
func (jas *JsonAssertion) IsMatchingStringArrayAt(path string, val []interface{}) (ok bool) {
	asserted := val
	val, ok = jas.arrayAtPath(splitPath(path))

	arraysMatch := func() (matches bool) {
		if len(val) == len(asserted) {
			itemMap := make(map[string]bool,len(val))
			matches = true
			for _, item := range val {  // range over found array, populating a map
				value := item.(string)   // todo: this will panic if fails. need gentler approach for testings
				itemMap[value]=true
			}
			for _, item := range asserted {  // range over found array, populating a map
				value := item.(string)   // todo: this will panic if fails. need gentler approach for testings
				if !itemMap[value] {
					matches = false
					break
				} // else asserted item was found in map
			}
		}
		return
	}
	// val comes from json retrieval, asserted comes from 'client' code
	return ok && arraysMatch()
}



