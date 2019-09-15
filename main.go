package main

import (
	// default go
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

const (
	// StringNewLine is an exported variable to declare string new line
	StringNewLine string = "\n"
	// StringSpace is an exported variable to declare string space
	StringSpace string = " "

	// CreateParkingLot is a command to execute create parking lot
	CreateParkingLot string = "create_parking_lot"
	// ParkVehicle is a command to execute to park a vehicle in parking lot
	ParkVehicle string = "park"
	// LeaveParking is a command to execute leave vehicle from parking lot
	LeaveParking string = "leave"
	// SlotNumberCarByColour is a command to execute slot number car by its register colour
	SlotNumberCarByColour string = "slot_numbers_for_cars_with_colour"
	// StatusParking is a command to execute parking lot status
	StatusParking string = "status"
	// SlotNumberForRegistrationNumber is a command to execute slot number registration
	SlotNumberForRegistrationNumber string = "slot_number_for_registration_number"
	// RegistrationNumbersForCarsWithColour is a command to see all vehicleID registration
	RegistrationNumbersForCarsWithColour string = "registration_numbers_for_cars_with_colour"
	// ExitCommand is a command to exit program
	ExitCommand string = "exit"

	// defaultNumber is a common default integer
	defaultNumber int = 0
)

// Vehicle this is struct to store vehicle information
type Vehicle struct {
	ID        int
	VehicleID string
	Color     string
}

func main() {

	// global variable to parking lot
	var parkingLot []*Vehicle
	// parkingLotAllocation size of created parkingLot and to compare when park command to append parking lot
	var parkingLotAllocation int

	inputFile, err := ioutil.ReadFile("functional_spec/fixtures/file_input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err, " \n")
	}

	// convert content to a 'string'
	stringifyInput := ConvertString(inputFile)

	// declare stringifyInput as array of string command input
	inputCommand := SplitString(stringifyInput, StringNewLine)

	for _, v := range inputCommand {

		// stackCommand = append(stackCommand, v)
		splitCommand := SplitString(v, StringSpace)

		command := splitCommand[0]

		if command == CreateParkingLot {
			if splitCommand[1] != "" {

				index := ConvertStringToInt(splitCommand[1])
				parkingLotAllocation = index
				parkingLot = CreateSlice(parkingLotAllocation)
				fmt.Printf("Created a parking with %d slots\n", index)

			}

		} else if command == ParkVehicle {
			if splitCommand[1] != "" && splitCommand[2] != "" {

				parkingLot, _ = Park(parkingLot, splitCommand[1], splitCommand[2])
			}
		} else if command == LeaveParking {
			index := ConvertStringToInt(splitCommand[1])
			checkSize := Size(parkingLot)
			// we need to have constraint:
			// - if command is "leave" but len of parking lot is 0, then do nothing
			// - if len of parking lot greater than 0, check if index size is always smaller than parking lot size, and  eligible to remove slice of index - 1
			if checkSize > defaultNumber && index < checkSize {
				parkingLot = Remove(parkingLot, index)
			}
			fmt.Printf("Slot number %d is free\n", index)

		} else if command == SlotNumberCarByColour {
			// void
			GetSlotRegistrationByColor(parkingLot, splitCommand[1])
		} else if command == StatusParking {
			// void
			GetParkingStatus(parkingLot)
		} else if command == SlotNumberForRegistrationNumber {
			// void
			GetSlotRegistrationByNumber(parkingLot, splitCommand[1])
		} else if command == RegistrationNumbersForCarsWithColour {
			// void
			GetRegistrationNumbersForCarsWithColour(parkingLot, splitCommand[1])
		} else if command == ExitCommand {
			os.Exit(2)
		}
	}
}

// SplitString func to return array of string
func SplitString(str, splitType string) []string {
	return strings.Split(str, splitType)
}

// ConvertStringToInt func to convert string to int type
func ConvertStringToInt(str string) int {
	index, err := strconv.Atoi(str)
	if err != nil {
		fmt.Print(err, " \n")
	}
	return index
}

// ConvertString func to convert []byte to string
func ConvertString(input []byte) string {
	return string(input)
}

// Peek func to get top value from stack list
func Peek(stackList []*Vehicle) *Vehicle {
	// get len of stackList
	size := len(stackList)

	// get last index value
	return stackList[size-1]
}

// Size func to get top value from stack list
func Size(stackList []*Vehicle) int {
	// get len of stackList
	size := len(stackList)
	return size
}

// Remove func to remove value of stack list by its index
func Remove(stackList []*Vehicle, index int) []*Vehicle {
	// set index i to nil
	for i, _ := range stackList {
		if i == index-1 {
			stackList[i] = nil
		}
	}
	return stackList
}

// CreateSlice func to create slize
func CreateSlice(size int) []*Vehicle {
	var newList []*Vehicle
	if size > 0 {
		newList = make([]*Vehicle, size)
	}
	return newList
}

// GetParkingStatus func to get information parking lot
func GetParkingStatus(stackList []*Vehicle) {

	// initialize tabwriter
	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 2, 8, 0, '\t', 0)

	defer w.Flush()

	// template header
	fmt.Fprintf(w, "%s\t%s\t%s\t \n", "Slot No.", "Registration No", "Colour")

	for _, v := range stackList {
		if v != nil {
			fmt.Fprintf(w, " %d\t%s\t%s\t \n", v.ID, v.VehicleID, v.Color)
		}
	}
}

// GetSlotRegistrationByColor func to get information of registration vehicle by its color category
func GetSlotRegistrationByColor(stackList []*Vehicle, color string) {
	// new list to append by its color

	newList := CreateSlice(defaultNumber)

	for _, v := range stackList {
		if v.Color == color {
			vehicle := &Vehicle{
				ID:        v.ID,
				VehicleID: v.VehicleID,
				Color:     v.Color,
			}
			newList = append(newList, vehicle)
		}
	}

	// get size new list
	size := Size(newList)
	if size > 0 {
		for i, v := range newList {
			if i < size-1 {
				fmt.Printf("%d,", v.ID)
			} else {
				fmt.Printf("%d\n", v.ID)
			}
		}
	}
}

// assume that registration number can have more than 1 and not unique
// GetSlotRegistrationByNumber func to get information of registration vehicle by its vehicleID category
func GetSlotRegistrationByNumber(stackList []*Vehicle, number string) {

	// this will check is registration number is exists
	var isVehicleIDFound bool

	newList := CreateSlice(defaultNumber)
	for _, v := range stackList {
		if v.VehicleID == number {
			isVehicleIDFound = true
			vehicle := &Vehicle{
				ID:        v.ID,
				VehicleID: v.VehicleID,
				Color:     v.Color,
			}
			newList = append(newList, vehicle)
		}
	}

	if isVehicleIDFound {
		// get size new list
		size := Size(newList)
		if size > 0 {
			for i, v := range newList {
				if i < size-1 {
					fmt.Printf("%d,", v.ID)
				} else {
					fmt.Printf("%d\n", v.ID)
				}
			}
		}
	} else {
		fmt.Printf("Not found\n")
	}

}

// GetRegistrationNumbersForCarsWithColour func to get information of registration vehicle by its color category
func GetRegistrationNumbersForCarsWithColour(stackList []*Vehicle, color string) {
	newList := CreateSlice(defaultNumber)

	for _, v := range stackList {
		if v.Color == color {
			vehicle := &Vehicle{
				ID:        v.ID,
				VehicleID: v.VehicleID,
				Color:     v.Color,
			}
			newList = append(newList, vehicle)
		}
	}

	// get size new list
	size := Size(newList)
	if size > 0 {
		for i, v := range newList {
			if i < size-1 {
				fmt.Printf("%s,", v.VehicleID)
			} else {
				fmt.Printf("%s\n", v.VehicleID)
			}
		}
	}
}

// Park func to insert park vehicle to parking lot if slot available
func Park(stackList []*Vehicle, vehicleID string, color string) ([]*Vehicle, int) {

	// allocatedPark. By default its value 0
	var allocatedPark int
	// isEligibleInsertNewList check if there is nil in new list
	var isEligibleInsertNewList bool

	// testing dereferenced
	for i, v := range stackList {
		// if there is nil in newList, append it.
		if v == nil {
			// initialize new vehicle to check if new list eligible to replace stack list
			// as long as v is nil, then set isEligibleInsertNewList to true
			// then break
			isEligibleInsertNewList = true

			vehicle := &Vehicle{
				ID:        (i + 1),
				VehicleID: vehicleID,
				Color:     color,
			}
			allocatedPark = vehicle.ID
			stackList[i] = vehicle
			break
		}
	}

	// assume that when array is full but admin wants to park new vehicle
	if isEligibleInsertNewList {
		// size of allocated
		fmt.Printf("Allocated slot number: %d \n", allocatedPark)
		return stackList, allocatedPark
	} else {
		fmt.Printf("Sorry, parking lot is full \n")
		return stackList, allocatedPark
	}
}
