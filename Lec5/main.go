package main

import (
	"fmt"
	"math"
)

func main() {
	// var value string

	// fmt.Scan(&value)

	// if strings.Contains(value, "чат") {
	// 	fmt.Println("БОТ")
	// } else {
	// 	fmt.Println("НЕ БОТ")
	// }

	// var (
	// 	one   string
	// 	two   string
	// 	three string
	// )

	// fmt.Scan(&one, &two, &three)

	// if one == "раз" || one == "один" && two == "два" && three == "три" {
	// 	fmt.Println("ОК")
	// 	return
	// }
	// if one == "1" && two == "2" && three == "3" {
	// 	fmt.Println("ОК")
	// 	return
	// }
	// fmt.Println("НЕ ПРАВИЛЬНО")

	// var (
	// 	text  string
	// 	coast float64
	// )

	// fmt.Scan(&text)
	// numberOfLetters := utf8.RuneCountInString(text)
	// coast = float64(numberOfLetters) * 0.23
	// intpart, fractpart := math.Modf(coast)
	// if intpart == 0 {
	// 	fmt.Printf("%.0f коп.\n", fractpart*100)
	// 	return
	// }
	// fmt.Printf("%.0f р. %.0f коп.\n", intpart, fractpart*100)

	var (
		a float64
		b float64
		c float64
		D float64
		// x float64
	)

	fmt.Scan(&a, &b, &c)
	if a == 0 {
		if b == 0 {
			fmt.Println("корней нет")
			return
		}
		fmt.Println("один корень")
		return
	}

	D = math.Pow(b, 2) - 4*a*c

	if D < 0 {
		fmt.Println("корней нет")
		return
	}
	if D == 0 {
		fmt.Println("один корень")
		return
	}
	if D > 0 {
		fmt.Println("два корня")
		return
	}

}
