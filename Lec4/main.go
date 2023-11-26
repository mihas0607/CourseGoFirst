package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	var (
		login string
		email string
	)

	fmt.Scan(&login)

	if utf8.RuneCountInString(login) >= 10 && !strings.Contains(login, "@") {
		fmt.Scan(&email)
		if strings.Contains(email, "@") && strings.Contains(email, ".") {
			fmt.Println("ОК")
		} else {
			fmt.Println("Некорректная почта")
		}

	} else {
		fmt.Println("Некорректный логин")
	}

	// var (
	// 	x1 float64
	// 	y1 float64
	// 	x2 float64
	// 	y2 float64
	// )

	// fmt.Scan(&x1, &y1, &x2, &y2)

	// if x1 < 9 && y1 < 9 && x2 < 9 && y2 < 9 {
	// 	if math.Abs(x2-x1) == 1 && math.Abs(y2-y1) == 2 {
	// 		fmt.Println("ДА")
	// 	} else if math.Abs(x2-x1) == 2 && math.Abs(y2-y1) == 1 {
	// 		fmt.Println("ДА")
	// 	} else {
	// 		fmt.Println("НЕТ")
	// 	}
	// } else {
	// 	fmt.Println("НЕТ")
	// }

}
