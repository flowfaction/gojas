// validates unknown bytes as a correct json document - mostly through built-in decoding error generation and checking
// provides higher level validations on the structure of the json documents, but only where amplifying value of built in checking
// executes assertions against the tree through a text dsl, or variadics, resty url equivalents, etc
// takes care of hairy reflection tasks to provide high-value code generation, dsl, _ ...

package gojas

import (
	"encoding/json"
	"strings"
	log "github.com/Sirupsen/logrus"
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
		log.Errorf("decoding error %v", err.Error())
	}
	return
}

func (jas *JsonAssertion) IsObjectAt(path string) (ok bool) {
	_, ok = jas.objectExists(splitPath(path), jas.receptacle)
	return
}

func (jas *JsonAssertion) IsNumberAt(path string, val float64) (ok bool) {
	asserted := val
	val, ok = jas.floatExists(splitPath(path))
	return ok && val == asserted
}

func (jas *JsonAssertion) IsBoolAt(path string, val bool) (ok bool) {
	asserted := val
	val, ok = jas.boolExists(splitPath(path))
	return ok && val == asserted
}

func (jas *JsonAssertion) IsStringAt(path string, val string) (ok bool) {
	asserted := val
	val, ok = jas.stringExists(splitPath(path))
	return ok && val == asserted
}

func (jas *JsonAssertion) IsFloatArrayAt(path string, val []interface{}) (ok bool) {
	asserted := val
	val, ok = jas.arrayExists(splitPath(path))
	//	log.Debugf(`
	//		val(%v)
	//		/val/type(%v)
	//		/exists(%v)
	//		/path(%v)
	//		/asserted(%v)
	//		/asserted/type(%v)
	//		/deepequal?(%v)`,val,reflect.TypeOf(val),ok,path,asserted,reflect.TypeOf(asserted),reflect.DeepEqual(val, asserted))

	//	for v := range val {
	//		 log.Debugf("item type in val/[]interface{} = (%v)",reflect.TypeOf(v))
	//	}
	//	for x := range asserted {
	//		 log.Debugf("item type in asserted/[]interface{} = (%v)",reflect.TypeOf(x))
	//	}

	arraysAreIdentical := func() (ident bool) {
		if len(val) == len(asserted) {
			ident = true
			for i, item := range val {
				aval := item.(float64)
				//				log.Debugf("a[item]=(%v)",aint)
				bval := asserted[i].(float64)
				//				log.Debugf("b[item]=(%v)",bint)
				if aval != bval {
					ident = false
					break
				}
			}
		}
		return
	}
	// val comes from json retrieval, asserted comes from 'client' code
	return ok && arraysAreIdentical()
}

func (jas *JsonAssertion) IsStringArrayAt(path string, val []interface{}) (ok bool) {
	asserted := val
	val, ok = jas.arrayExists(splitPath(path))

	arraysAreIdentical := func() (ident bool) {
		if len(val) == len(asserted) {
			ident = true
			for i, item := range val {
				aval := item.(string)   // todo: this will panic if fails. need gentler approach for testings
				//				log.Debugf("a[item]=(%v)",aint)
				bval := asserted[i].(string)
				//				log.Debugf("b[item]=(%v)",bint)
				if strings.Compare(aval, bval)!=0 {
					ident = false
					break
				}
			}
		}
		return
	}
	// val comes from json retrieval, asserted comes from 'client' code
	return ok && arraysAreIdentical()
}

