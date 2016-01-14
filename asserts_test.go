package gojas

import (
	"testing"
)


func TestAssertObject(t *testing.T) {
	path := "/user/properties/object/innerObject"

	if !AssertObjectAtPath(t,asserted_json_data,path) {
		t.Fatal("Failed to assert existence of known object at known path")
	}

	if AssertObjectAtPath(t,asserted_json_data,"/user/properties/object/inner__"){
		t.Fatal("Failed to assert non-existence of object at fake or bad path")
	}

	sim := *t
	_ = AssertObjectAtPath(&sim,bad_json_data,path)
	if !sim.Failed() { // should have marked test as failed
		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
	}


}


//Tests whether an object exists at a path, and then tests that it contains known property names aka keys.
func TestAssertObjectKeys(t *testing.T) {
	path := "/user/properties/object/innerObject"
	bad_path := "/user/properties/object/inner_"
	keys := []string{"foo","baz","key"}

	if !AssertObjectAtPathWithKeys(t,asserted_json_data,path,keys) {
		t.Fatal("Failed to assert existence of object at path with known keys.")
	}

	sim := *t
	bad_path_keys := AssertObjectAtPathWithKeys(&sim,asserted_json_data, bad_path,keys)
	if !sim.Failed() || bad_path_keys {
		t.Fatal("Failed to assert non-existence of object at fake or bad path, failed to fail the test.")
	} else {
		t.Log("Test assertion successfully detected non-existent path, and failed the test object. win!")
	}


	// assert that it fails when it should
	sim = *t
	keys_ok := AssertObjectAtPathWithKeys(&sim,asserted_json_data,path,[]string{"foo","baz","ishouldnotexist"})
	if !sim.Failed() || keys_ok { // should have marked test as failed, but didn't
		t.Fatal("Test assertion failed to detect missing property key in target JSON object. fail!")
	} else {
		t.Log("Test assertion successfully detected non-existent but asserted key in json doc. win!")
	}

	//assert that it fails when it can't parse the json data to begin with
	sim = *t
	_ = AssertObjectAtPathWithKeys(&sim,bad_json_data,path,keys)
	if !sim.Failed() { // should have marked test as failed
		t.Fatal("Failed to fail test when detecting parse error or failing to MakeJsonAssertion()")
	} else {
		t.Log("test assertion had error parsing intentionally bad json data, as expected. win!")
	}


}


func TestAssertNumber(t *testing.T) {
	var val float64 = 11235
	path := "/user/properties/object/innerObject/baz"

	if !AssertNumberAtPath(t,asserted_json_data,path,val) {
		t.Fatal("Failed to assert matching float value at given path")
	}

	if AssertNumberAtPath(t,asserted_json_data,path,val*2){
		t.Fatal("Failed to fail assert of non-matching float value at known path")
	}

	sim := *t
	_ = AssertNumberAtPath(&sim,bad_json_data,path,val)
	if !sim.Failed() { // should have marked test as failed
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

	if !AssertFloatArrayAtPath(t,asserted_json_data,path,val) {
		t.Fatal("Failed to pass asserted float array at known path")
	}

	val = append(val,9.0)
	if AssertFloatArrayAtPath(t,asserted_json_data,path,val){
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
	if !AssertStringArrayAtPath(t,asserted_json_data,path,val) {
		t.Fatal("Failed to pass asserted string array at known path")
	}


	val = append(val,"9")
	if AssertStringArrayAtPath(t,asserted_json_data,path,val) {
		t.Fatal("Failed to fail assert of modified string array at known path")
	}

	sim := *t
	_ = AssertStringArrayAtPath(&sim,bad_json_data,path,val)
	if !sim.Failed() { // should have marked test as failed
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
	if !AssertStringAtPath(t,asserted_json_data,path,val) {
		t.Fatal("Failed to pass assert string('foobar') at known path")
	}

	if AssertStringAtPath(t,asserted_json_data,path,"shouldnotbethere") {
		t.Fatal("Failed to fail assert of string('shouldnotbethere') at known path")
	}


	sim := *t
	_ = AssertStringAtPath(&sim,bad_json_data,path,val)
	if !sim.Failed() { // should have marked test as failed
		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
//		// it did fail, but we want to reset the status for the purposes of this outer self-test
		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
	}

}

func TestAssertBool(t *testing.T) {
	path := "/user/properties/boolean/value"

	if !AssertBoolAtPath(t,asserted_json_data,path,true) {
		t.Fatal("Failed to pass assert boolean(true) at known path")
	}

	if AssertBoolAtPath(t,asserted_json_data,path,false) {
		t.Fatal("Failed to fail assert of not boolean(false) at known path")
	}

	sim := *t
	_ = AssertBoolAtPath(&sim,bad_json_data,path,true)
	if !sim.Failed() { // should have marked test as failed
		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
//		// it did fail, but we want to reset the status for the purposes of this outer self-test
		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
	}


}

