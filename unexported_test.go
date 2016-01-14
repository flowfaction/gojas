package gojas

import (
	"testing"
)

func TestIdenticalStringSlices(t *testing.T) {
	sliceLeft := []interface{}{"1", "1", "2", "3", "5", "8", "13"}
	sliceRight := []interface{}{"1", "1", "2", "3", "5", "8", "13"}

	left, right := sliceLeft, sliceRight

	identical := areIdenticalStringInterfaceSlices(left, right)
	if !identical {
		t.Errorf("Failed to detect two identical string slices, as interface slices.")
	}

	left, right = []interface{}{"1", 1, "2", "3", "5", "8", "13"}, sliceRight

	identical = areIdenticalStringInterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a left-val element that is not a string, in interface slices.")
	}

	left, right = sliceLeft, []interface{}{"1", "1", "2", "3", "5", "8", 13}
	identical = areIdenticalStringInterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a right-val element that is not a string, in interface slices.")
	}

}

func TestMatchingStringSlices(t *testing.T) {
	sliceLeft  := []interface{}{"1", "1", "2", "3", "5", "8", "13"}
	sliceRight := []interface{}{"13", "1", "1", "2", "3", "5", "8"}

	left, right := sliceLeft, sliceRight

	// should match even if items are out of order
	identical := areMatchingStringInterfaceSlices(left, right)
	if !identical {
		t.Errorf("Failed to detect two identical string slices, as interface slices.")
	}

	left, right = []interface{}{"1", 1, "2", "3", "5", "8", "13"}, sliceRight
	identical = areMatchingStringInterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a left-val element that is not a string, in interface slices.")
	}

	left, right = sliceLeft, []interface{}{"1", "1", "2", "3", "5", "8", 13}
	identical = areMatchingStringInterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a right-val element that is not a string, in interface slices.")
	}

	left, right = []interface{}{"42", "1", "2", "3", "5", "8", "13"}, sliceRight
	// should not match because the left starts with a "42"
	identical = areMatchingStringInterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a right-val element that doesn't appear in the left slice")
	}


	left, right = sliceLeft, []interface{}{"1", "1", "1", "2", "3", "5", "8","13"}
	// should not match because the right has 3 '1's and left only has 2
	identical = areMatchingStringInterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a right-val element that matches a left-val, but at different frequency")
	}



}

func TestIdenticalFloatSlices(t *testing.T) {

	sliceLeft := []interface{}{1.1, 1.1, 2.1, 3.1, 5.1, 8.1, 13.1}
	sliceRight := []interface{}{1.1, 1.1, 2.1, 3.1, 5.1, 8.1, 13.1}

	left, right := sliceLeft, sliceRight

	identical := areIdenticalFloat64InterfaceSlices(left, right)
	if !identical {
		t.Errorf("Failed to detect two identical float slices, as interface slices.")
	}

	left, right = []interface{}{1.1, "1.1", 2.1, 3.1, 5.1, 8.1, 13.1}, sliceRight
	identical = areIdenticalFloat64InterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a left-val element that is not a float, in interface slices.")
	}

	left, right = sliceLeft, []interface{}{1.1, "1.1", 2.1, 3.1, 5.1, 8.1, 13.1}

	identical = areIdenticalFloat64InterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a right-val element that is not a float, in interface slices.")
	}

}
