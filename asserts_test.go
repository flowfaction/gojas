package gojas

import (
	"testing"
)


func TestAssertObject(t *testing.T) {
	path := "/user/properties/object/innerObject"

	ok := AssertObjectAtPath(t,asserted_json_data,path)
	if !ok {
		t.Fatal("Failed to assert existence of known object at known path")
	}

	ok = AssertObjectAtPath(t,asserted_json_data,"/user/properties/object/inner__")
	if ok {
		t.Fatal("Failed to assert non-existence of object at fake or bad path")
	}

	cloned_t := *t
	_ = AssertObjectAtPath(&cloned_t,bad_json_data,path)
	if !cloned_t.Failed() { // should have marked test as failed
		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
//		// it did fail, but we want to reset the status for the purposes of this outer self-test
		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
	}


}


func TestAssertNumber(t *testing.T) {

	var val float64 = 11235
	path := "/user/properties/object/innerObject/baz"

	ok := AssertNumberAtPath(t,asserted_json_data,path,val)
	if !ok {
		t.Fatal("Failed to assert matching float value at given path")
	}

	ok = AssertNumberAtPath(t,asserted_json_data,path,val*2)
	if ok {
		t.Fatal("Failed to fail assert of non-matching float value at known path")
	}

	cloned_t := *t
	_ = AssertNumberAtPath(&cloned_t,bad_json_data,path,val)
	if !cloned_t.Failed() { // should have marked test as failed
		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
//		// it did fail, but we want to reset the status for the purposes of this outer self-test
		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
	}

}

func TestAssertFloatArray(t *testing.T) {

	original_val := []interface{}{1.0, 1.0, 2.0, 3.0, 5.0, 8.0}

	path := "/user/properties/numberArray/value"
	val := original_val

	ok := AssertFloatArrayAtPath(t,asserted_json_data,path,val)
	if !ok {
		t.Fatal("Failed to pass asserted float array at known path")
	}

	val = append(val,9.0)
	ok = AssertFloatArrayAtPath(t,asserted_json_data,path,val)
	if ok {
		t.Fatal("Failed to fail assert of modified float array at known path")
	}

	cloned_t := *t
	_ = AssertFloatArrayAtPath(&cloned_t,bad_json_data,path,val)
	if !cloned_t.Failed() { // should have marked test as failed
		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
//		// it did fail, but we want to reset the status for the purposes of this outer self-test
		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
	}


}

func TestAssertStringArray(t *testing.T) {

	original_val := []interface{}{"1","1","2","3","5","8"}
	val := original_val

	path := "/user/properties/stringArray/value"
	ok := AssertStringArrayAtPath(t,asserted_json_data,path,val)
	if !ok {
		t.Fatal("Failed to pass asserted string array at known path")
	}


	val = append(val,"9")
	ok = AssertStringArrayAtPath(t,asserted_json_data,path,val)
	if ok {
		t.Fatal("Failed to fail assert of modified string array at known path")
	}

	cloned_t := *t
	_ = AssertStringArrayAtPath(&cloned_t,bad_json_data,path,val)
	if !cloned_t.Failed() { // should have marked test as failed
		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
//		// it did fail, but we want to reset the status for the purposes of this outer self-test
		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
	}



}

func TestAssertString(t *testing.T) {
	path := "/user/properties/string/value"

	val := "foobar"
	ok := AssertStringAtPath(t,asserted_json_data,path,val)
	if !ok {
		t.Fatal("Failed to pass assert string('foobar') at known path")
	}

	ok = AssertStringAtPath(t,asserted_json_data,path,"shouldnotbethere")
	if ok {
		t.Fatal("Failed to fail assert of string('shouldnotbethere') at known path")
	}


	cloned_t := *t
	_ = AssertStringAtPath(&cloned_t,bad_json_data,path,val)
	if !cloned_t.Failed() { // should have marked test as failed
		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
//		// it did fail, but we want to reset the status for the purposes of this outer self-test
		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
	}

}

func TestAssertBool(t *testing.T) {
	path := "/user/properties/boolean/value"

	ok := AssertBoolAtPath(t,asserted_json_data,path,true)
	if !ok {
		t.Fatal("Failed to pass assert boolean(true) at known path")
	}

	ok = AssertBoolAtPath(t,asserted_json_data,path,false)
	if ok {
		t.Fatal("Failed to fail assert of not boolean(false) at known path")
	}

	cloned_t := *t
	_ = AssertBoolAtPath(&cloned_t,bad_json_data,path,true)
	if !cloned_t.Failed() { // should have marked test as failed
		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
//		// it did fail, but we want to reset the status for the purposes of this outer self-test
		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
	}


}

