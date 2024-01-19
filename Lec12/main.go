package main

import "fmt"

func main() {
	// var (
	// 	width int
	// 	height int
	// 	symbol string

	// )

	// fmt.Scan(&width, &height, &symbol)

	// for i:=0; i < height; i++ {
	// 	for j:=0; j < width; j++ {
	// 		if i == 0 || i == height-1 {
	// 			fmt.Print(symbol)
	// 		} else {
	// 			fmt.Print(symbol)
	// 			for k := 0; k < width-2; k++ {
	// 				fmt.Print(" ")
	// 			}
	// 			fmt.Print(symbol)
	// 			break
	// 		}

	// 	}
	// 	fmt.Println()
	// }

	var (
		sum   int
		count int
	)
	calf := 5
	// cow := 2 * calf
	// bull := 4 * calf
	fmt.Scan(&sum, &count)
	countCalf := sum / calf
	fmt.Println(countCalf)

}
