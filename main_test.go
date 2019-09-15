package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	inputCommand := []string{"create_parking_lot 6", "park KA-01-HH-1234 White",
		"park KA-01-HH-9999 White",
		"park KA-01-BB-0001 Black",
		"park KA-01-HH-7777 Red",
		"park KA-01-HH-2701 Blue",
		"park KA-01-HH-3141 Black",
		"leave 4",
		"status",
		"park KA-01-P-333 White",
		"park DL-12-AA-9999 White",
		"registration_numbers_for_cars_with_colour White",
		"slot_numbers_for_cars_with_colour White",
		"slot_number_for_registration_number KA-01-HH-3141",
		"slot_number_for_registration_number MH-04-AY-1111",
	}

	ProcessInput(inputCommand)
}

func TestLeaveWithoutError(t *testing.T) {
	stackList := make([]*Vehicle, 6)
	removeID := 2
	exists := true

	for i, v := range stackList {

		if v == nil {
			vehicle := &Vehicle{

				ID:        (i + 1),
				VehicleID: "KA-01-HH-1234",
				Color:     "White",
			}
			stackList = append(stackList, vehicle)
		}
	}

	// error log
	expected := ""
	err := "vehicleID exists"

	actual := Remove(stackList, removeID)
	list := Remove(stackList, removeID)
	for i, _ := range list {
		if i == removeID-1 {
			list[i] = nil
			exists = false
		}
	}
	if actual == nil && exists == false {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, err)
	}
}

func TestCreateParkingLotWithoutError(t *testing.T) {
	actualAllocated := 2
	actualStackList := make([]*Vehicle, actualAllocated)

	stackList := CreateSlice(actualAllocated)

	if len(stackList) != len(actualStackList) {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", "expected length of stackList", "different length of Vehicle")
	}

}

func TestParkWithoutError(t *testing.T) {
	stackList := make([]*Vehicle, 6)
	exists := false

	for i, v := range stackList {

		if v == nil {

			// dummy index 4 is nil
			// this will  bypass if there is nil in stackList, then will insert in Park()
			if i != 3 {
				vehicle := &Vehicle{

					ID:        (i + 1),
					VehicleID: "KA-01-HH-1234",
					Color:     "White",
				}
				stackList = append(stackList, vehicle)
			}
		}
	}

	// error log
	expected := ""
	err := "vehicleID exists"

	actualAllocated := 3

	// actual, actualAllocated := Park(stackList, "KA-01-HH-3141", "Black")
	list, allocated := Park(stackList, "KA-01-HH-3141", "Black")
	for i, v := range list {
		if i == allocated {
			if v == nil {
				exists = true
			}
		}
	}
	if actualAllocated != allocated {
		if allocated != 3 && exists == false {
			t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, err)
		}
	}

}
