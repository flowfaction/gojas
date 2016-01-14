package gojas

import (
	"testing"
)

func TestIdenticalStringSlices(t *testing.T) {
	sliceLeft := []interface{}{"1", "1", "2", "3", "5", "8", "13"}
	sliceRight := []interface{}{"1", "1", "2", "3", "5", "8", "13"}

	left, right := sliceLeft, sliceRight

	if !areIdenticalStringInterfaceSlices(left, right) {
		t.Errorf("Failed to detect two identical string slices, as interface slices.")
	}

	left, right = []interface{}{"1", 1, "2", "3", "5", "8", "13"}, sliceRight

	if areIdenticalStringInterfaceSlices(left, right) {
		t.Errorf("Failed to detect a left-val element that is not a string, in interface slices.")
	}

	left, right = sliceLeft, []interface{}{"1", "1", "2", "3", "5", "8", 13}
	if areIdenticalStringInterfaceSlices(left, right) {
		t.Errorf("Failed to detect a right-val element that is not a string, in interface slices.")
	}

}

func TestMatchingStringSlices(t *testing.T) {
	sliceLeft  := []interface{}{"1", "1", "2", "3", "5", "8", "13"}
	sliceRight := []interface{}{"13", "1", "1", "2", "3", "5", "8"}

	left, right := sliceLeft, sliceRight

	// should match even if items are out of order
	if !areMatchingStringInterfaceSlices(left, right) {
		t.Errorf("Failed to detect two identical string slices, as interface slices.")
	}

	//detect non-string element on left
	left, right = []interface{}{"1", 1, "2", "3", "5", "8", "13"}, sliceRight
	if areMatchingStringInterfaceSlices(left, right) {
		t.Errorf("Failed to detect a left-val element that is not a string, in interface slices.")
	}

	//detect non-string element on right
	left, right = sliceLeft, []interface{}{"1", "1", "2", "3", "5", "8", 13}
	if areMatchingStringInterfaceSlices(left, right) {
		t.Errorf("Failed to detect a right-val element that is not a string, in interface slices.")
	}

	// detect key in left that is not in right
	left, right = []interface{}{"11", "1", "2", "3", "5", "8", "13"}, sliceRight
	if areMatchingStringInterfaceSlices(left, right) {
		t.Errorf("Failed to detect a right-val element that doesn't appear in the left slice")
	}


	left, right = sliceLeft, []interface{}{"1", "1", "1", "2", "3", "5", "8","13"}
	// should not match because the right has 3 '1's and left only has 2
	if areMatchingStringInterfaceSlices(left, right) {
		t.Errorf("Failed to detect a right-val element that matches a left-val, but at different frequency")
	}



}


func TestIdenticalFloatSlices(t *testing.T) {

	sliceLeft := []interface{}{1.1, 1.1, 2.1, 3.1, 5.1, 8.1, 13.1}
	sliceRight := []interface{}{1.1, 1.1, 2.1, 3.1, 5.1, 8.1, 13.1}

	left, right := sliceLeft, sliceRight

	if !areIdenticalFloat64InterfaceSlices(left, right) {
		t.Errorf("Failed to detect two identical float slices, as interface slices.")
	}

	left, right = []interface{}{1.1, "1.1", 2.1, 3.1, 5.1, 8.1, 13.1}, sliceRight
	if areIdenticalFloat64InterfaceSlices(left, right) {
		t.Errorf("Failed to detect a left-val element that is not a float, in interface slices.")
	}

	left, right = sliceLeft, []interface{}{1.1, "1.1", 2.1, 3.1, 5.1, 8.1, 13.1}

	if areIdenticalFloat64InterfaceSlices(left, right) {
		t.Errorf("Failed to detect a right-val element that is not a float, in interface slices.")
	}

}

func TestMatchingFloadSlices(t *testing.T) {
	sliceLeft  := []interface{}{1.1, 1.1, 2.1, 3.1, 5.1, 8.1, 13.1}
	sliceRight := []interface{}{1.1, 1.1, 2.1, 3.1, 5.1, 8.1, 13.1}

	left, right := sliceLeft, sliceRight

	// should match even if items are out of order
	if !areMatchingFloat64InterfaceSlices(left, right) {
		t.Errorf("Failed to detect two identical float64 slices, as interface slices.")
	}

	//detect non-float element on left
	left, right = []interface{}{1.1, "1.1", 2.1, 3.1, 5.1, 8.1, 13.1}, sliceRight
	if areMatchingFloat64InterfaceSlices(left, right) {
		t.Errorf("Failed to detect a left-val element that is not a float64, in interface slices.")
	}

	//detect non-float element on right
	left, right = sliceLeft, []interface{}{1.1, 1.1, "2.1", 3.1, 5.1, 8.1, 13.1}
	if areMatchingFloat64InterfaceSlices(left, right) {
		t.Errorf("Failed to detect a right-val element that is not a float64, in interface slices.")
	}

	// detect key in left that is not in right
	left, right = []interface{}{1.1, 1.7, 2.1, 3.1, 5.1, 8.1, 13.1}, sliceRight
	if areMatchingFloat64InterfaceSlices(left, right) {
		t.Errorf("Failed to detect a left-val key that doesn't appear in the right slice")
	}

	// detect key in right that is not in left
	left,right = sliceLeft, []interface{}{1.1, 1.1, 2.7, 3.1, 5.1, 8.1, 13.1}
	if areMatchingFloat64InterfaceSlices(left, right) {
		t.Errorf("Failed to detect a right-val key that doesn't appear in the left slice")
	}


}
