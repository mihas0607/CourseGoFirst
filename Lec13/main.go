package main

import "fmt"

func main() {
	var variable int = 30
	var zeroPointer *int //zeroValue имеет значение nil (это указатель в никуда)
	fmt.Printf("Type %T and value %v\n", zeroPointer, zeroPointer)
	if zeroPointer == nil {
		zeroPointer = &variable
		fmt.Printf("After initializatoin type %T and value %v\n", zeroPointer, zeroPointer)
	}
}

func CalculateYearsNew(years int) (result [3]int) {
	var catYears, dogYears int
	for i := 1; i <= years; i++ {
		switch {
		case i == 1:
			catYears += 15
			dogYears += 15
		case i == 2:
			catYears += 9
			dogYears += 9
		default:
			catYears += 4
			dogYears += 5
		}		
	}
	result[0] = years
	result[1] = catYears
	result[2] = dogYears
	return
}

func CalculateYears(years int) (result [3]int) {
	var catYears, dogYears int
	for i := 1; i <= years; i++ {
		if i == 1 {
			catYears += 15
			dogYears += 15
		} else if i == 2 {
			catYears += 9
			dogYears += 9
		} else {
			catYears += 4
			dogYears += 5
		}
	}
	result[0] = years
	result[1] = catYears
	result[2] = dogYears
	return
}

func CalculateYearsBest(years int) (result [3]int) {
switch years {
case 1: result = [3]int{1, 15, 15}
case 2: result = [3]int{2, 24, 24}
default: result = [3]int{years, 24 + 4 * (years - 2), 24 + 5 * (years - 2)}
}
return
}

func Between(a, b int) []int {
	slice := []int{}
	for i := a; i <= b; i++ {
		slice = append(slice, i)
	}
	return slice
}
