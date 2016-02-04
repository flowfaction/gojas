package gojas

import (
	"testing"
)

const LogPasses = false


func assertSimFailed(sim testing.T, t *testing.T) {
	if sim.Failed() {
		if LogPasses{t.Log("Assertion set T to failed. PASS")}
	} else {
		t.Fatal("AssertObjectAtPath:Assertion did not set T to failed. FAIL")
	}
}

func TestAssertShortPath(t *testing.T) {
	path := "/user"

	if AssertObjectAtPath(t,asserted_json_data,path) {
		if LogPasses {t.Log("Found object at short path. PASS")}
	} else {
		t.Fatal("AssertObjectAtPath-Failed to assert existence of known object at known path")
	}

	sim := *t
	if AssertObjectAtPath(&sim,asserted_json_data,"/nope") {
		t.Fatal("AssertObjectAtPath-Failed to fail an assertion about an object at a non-existent path.")
	} else {
		assertSimFailed(sim,t)
	}

}


func TestAssertObject(t *testing.T) {
	path := "/user/properties/object/innerObject"

	if !AssertObjectAtPath(t,asserted_json_data,path) {
		t.Fatal("AssertObjectAtPath-Failed to assert existence of known object at known path")
	}

	sim := *t
	if AssertObjectAtPath(&sim,asserted_json_data,"/user/properties/object/idontexist"){
		t.Fatal("AssertObjectAtPath-Failed to assert non-existence of object at fake or bad path")
	} else {
		assertSimFailed(sim,t)
	}

	sim = *t
	if AssertObjectAtPath(&sim,bad_json_data,path) {
		t.Fatal("AssertObjectAtPath-Failed to fail test when detecting parse error or failing to make a JsonAssertion")
	} else {
		assertSimFailed(sim,t)
	}


}


//Tests whether an object exists at a path, and then tests that it contains known property names aka keys.
func TestAssertObjectKeys(t *testing.T) {
	path := "/user/properties/object/innerObject"
	bad_path := "/user/properties/object/idontexist"
	keys := []string{"foo","baz","key"}

	if !AssertObjectAtPathWithKeys(t,asserted_json_data,path,keys) {
		t.Fatal("AssertObjectAtPathWithKeys-Failed to assert existence of object at path with known keys.")
	}

	sim := *t
	if AssertObjectAtPathWithKeys(&sim,asserted_json_data, bad_path,keys) {
		t.Fatal("AssertObjectAtPathWithKeys-Failed to fail assertion against bad path")
	} else {
		assertSimFailed(sim,t)
	}


	sim = *t
	if AssertObjectAtPathWithKeys(&sim,asserted_json_data,path,[]string{"foo","baz","ishouldnotexist"}) {
		t.Fatal("AssertObjectAtPathWithKeys-Failed to fail assertion against bad values")
	} else {
		assertSimFailed(sim,t)
	}

	//assert that it fails when it can't parse the json data to begin with
	sim = *t
	if AssertObjectAtPathWithKeys(&sim,bad_json_data,path,keys) {
		t.Fatal("AssertObjectAtPathWithKeys-Failed to fail assertion against bad json data")
	} else {
		assertSimFailed(sim,t)
	}

}


func TestAssertNumber(t *testing.T) {
	var val float64 = 11235
	path := "/user/properties/object/innerObject/baz"

	if !AssertNumberAtPath(t,asserted_json_data,path,val) {
		t.Fatal("Failed to assert matching float value at given path")
	}

	sim := *t
	if AssertNumberAtPath(&sim,asserted_json_data,path,val*2){
		t.Fatal("Failed to fail assert of non-matching float value at known path")
	} else {
		assertSimFailed(sim,t)
	}

	sim = *t
	if AssertNumberAtPath(&sim,bad_json_data,path,val) {
		t.Fatal("AssertObjectAtPathWithKeys-Failed to fail assertion against bad json data")
	} else {
		assertSimFailed(sim,t)
	}


}

//func TestAssertFloatArray(t *testing.T) {
//	original_val := []interface{}{1.0, 1.0, 2.0, 3.0, 5.0, 8.0}
//
//	path := "/user/properties/numberArray/value"
//	val := original_val
//
//	if !AssertFloatArrayAtPath(t,asserted_json_data,path,val) {
//		t.Fatal("Failed to pass asserted float array at known path")
//	}
//
//	val = append(val,9.0)
//	if AssertFloatArrayAtPath(t,asserted_json_data,path,val){
//		t.Fatal("Failed to fail assert of modified float array at known path")
//	}
//
//	sim := clone(t)
//	_ = AssertFloatArrayAtPath(sim,bad_json_data,path,val)
//	if !sim.Failed() { // should have marked test as failed
//		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
//		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
//	} else {
////		// it did fail, but we want to reset the status for the purposes of this outer self-test
//		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
//	}
//
//}
//
//func TestAssertStringArray(t *testing.T) {
//
//	original_val := []interface{}{"1","1","2","3","5","8"}
//	val := original_val
//
//	path := "/user/properties/stringArray/value"
//	if !AssertStringArrayAtPath(t,asserted_json_data,path,val) {
//		t.Fatal("Failed to pass asserted string array at known path")
//	}
//
//
//	val = append(val,"9")
//	if AssertStringArrayAtPath(t,asserted_json_data,path,val) {
//		t.Fatal("Failed to fail assert of modified string array at known path")
//	}
//
//	sim := clone(t)
//	_ = AssertStringArrayAtPath(sim,bad_json_data,path,val)
//	if !sim.Failed() { // should have marked test as failed
//		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
//		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
//	} else {
////		// it did fail, but we want to reset the status for the purposes of this outer self-test
//		t.Logf("test assertion had error parsing intentionally bad json data, as	 expected. win!")
//	}
//
//
//
//}
//
//func TestAssertString(t *testing.T) {
//	path := "/user/properties/string/value"
//
//	val := "foobar"
//	if !AssertStringAtPath(t,asserted_json_data,path,val) {
//		t.Fatal("Failed to pass assert string('foobar') at known path")
//	}
//
//	if AssertStringAtPath(clone(t),asserted_json_data,path,"shouldnotbethere") {
//		t.Fatal("AssertStringAtPath Failed to mark testing.T as failed when not finding 'shouldnotbethere' at known path")
//	}
//
//	sim := clone(t)
//	if AssertStringAtPath(sim,bad_json_data,path,val) {
//		// bad. it should have returned false
//		t.Logf("AssertStringAtPath failed to detect parsing error on intentionally bad data. fail!")
//		t.Fatal("AssertStringAtPath failed to detect parse error or failed to make a JsonAssertion")
//	} else {
//		// but did it fail the T?
//		if sim.Failed() {
//			t.Logf("PASS: AssertStringAtPath correctly marked the testing.T as failed.")
//		} else {
//			t.Fatal("FAIL: AssertStringAtPath failed to mark testing.T as failed")
//		}
//	}
//
//}
//
//func TestAssertBool(t *testing.T) {
//	path := "/user/properties/boolean/value"
//
//	if !AssertBoolAtPath(t,asserted_json_data,path,true) {
//		t.Fatal("Failed to pass assert boolean(true) at known path")
//	}
//
//	if AssertBoolAtPath(t,asserted_json_data,path,false) {
//		t.Fatal("Failed to fail assert of not boolean(false) at known path")
//	}
//
//	sim := clone(t)
//	if AssertBoolAtPath(sim,bad_json_data,path,true) {
//		// the assert should have caught the bad data
//		t.Fatal("Failed to fail an assertion about a boolean[%v] at a non-existent path, in unparseable JSON.",true)
//	} else {
//		// the assert caught it - good - and returned false, but did it fail the 'T' ?
//
//	}
//
//
//
//	if !sim.Failed() { // should have marked test as failed
//		t.Logf("test assertion failed to detect parsing error on intentionally bad data. fail!")
//		t.Fatal("Failed to fail test when detecting parse error or failing to make a JsonAssertion")
//	} else {
////		// it did fail, but we want to reset the status for the purposes of this outer self-test
//		t.Logf("test assertion had error parsing intentionally bad json data, as expected. win!")
//	}
//
//
//}
//
