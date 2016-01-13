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
	sliceLeft := []interface{}{"1", "1", "2", "3", "5", "8", "13"}
	sliceRight := []interface{}{"13", "1", "1", "2", "3", "5", "8"}

	left, right := sliceLeft, sliceRight

	// should match even if items are out of order
	identical := areMatchingStringInterfaceSlices(left, right)
	if !identical {
		t.Errorf("Failed to detect two identical string slices, as interface slices.")
	}

	left, right = []interface{}{"1", 1, "2", "3", "5", "8", "13"}, sliceRight
	// should not match if it finds the numeric '1' instead of a string
	identical = areMatchingStringInterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a left-val element that is not a string, in interface slices.")
	}

	left, right = sliceLeft, []interface{}{"1", "1", "2", "3", "5", "8", 13}
	// should not match if it finds the numeric '13' instead of a string
	identical = areMatchingStringInterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a right-val element that is not a string, in interface slices.")
	}

	left, right = []interface{}{"42", "1", "2", "3", "5", "8", "13"}, sliceRight
	// should not match if it finds the numeric '13' instead of a string
	identical = areMatchingStringInterfaceSlices(left, right)
	if identical {
		t.Errorf("Failed to detect a right-val element that is not a string, in interface slices.")
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
