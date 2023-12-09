package main

import (
	"fmt"
)

func main() {

	// for i := 0; i <= 10; i++ {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("After loop")

	// outer:
	// for i := 0; i <= 2; i++ {
	// 	for j := 1; j <= 3; j++ {
	// 		fmt.Println(i,j)
	// 		break outer
	// 	}
	// }

	// i := 0

	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// 	var pass string
	// outer:
	// 	for {
	// 		fmt.Print("Enter a pass: ")
	// 		fmt.Scan(&pass)
	// 		if strings.Contains(pass, "1234") {
	// 			fmt.Println("Weak pass")
	// 		} else {
	// 			fmt.Println("Pass acepted")
	// 			break outer
	// 		}
	// 	}
	// var value string

	// for {
	// 	fmt.Scan(&value)
	// 	if len(value) != 0 {
	// 		fmt.Println(value)
	// 	} else {
	// 		fmt.Scanln(&value)
	// 		break
	// 	}
	// }

	// var value float64

	// for {
	// 	fmt.Scan(&value)
	// 	if value == 0 {
	// 		return
	// 	}
	// 	fmt.Println(value)
	// }

	// var (
	// 	pass1 string
	// 	pass2 string
	// )

	// for {
	// 	fmt.Scan(&pass1, &pass2)
	// 	if utf8.RuneCountInString(pass1) < 8 {
	// 		fmt.Println("Слишком короткий пароль!")
	// 	} else if strings.Contains(pass1, "123") || strings.Contains(pass1, "qwe") {
	// 		fmt.Println("Слишком простой пароль!")
	// 	} else if pass1 != pass2 {
	// 		fmt.Println("Введенные пароли различаются!")
	// 	} else {
	// 		fmt.Println("Ну наконец-то!")
	// 		return
	// 	}
	// }

	// var (
	// 	puls  float64
	// 	count int = 0
	// 	pulsMin float64 = 140
	// 	pulsMax float64
	// )

	// for {
	// 	fmt.Scan(&puls)
	// 	if puls >= 100 && puls <= 140 {
	// 		count++
	// 		if puls > pulsMax {
	// 			pulsMax = puls
	// 		}
	// 		if puls < pulsMin {
	// 			pulsMin = puls
	// 		}
	// 	}
	// 	if puls < 0 {
	// 		fmt.Println(count)
	// 		fmt.Printf("%.1f %.1f", pulsMin, pulsMax)
	// 		return
	// 	}
	// }

	// var (
	// 	count  int
	// 	number int
	// 	sum    int
	// )
	// fmt.Scan(&count)
	// for i := 0; i < count; i++ {
	// 	fmt.Scan(&number)
	// 	if i == 0 {
	// 		sum = number
	// 	} else if i%2 != 0 {
	// 		sum = sum - number
	// 	} else {
	// 		sum = sum + number
	// 	}
	// }
	// fmt.Println(sum)

	// for i := 0; i <= 20; i++ {
	// 	if i > 10 && i <= 15 {
	// 		continue
	// 	}
	// 	fmt.Printf("Current value: %d\n", i)
	// }
	// fmt.Println("After for loop with CONTINUE")

	// var (
	// 	count        int
	// 	numerator    int
	// 	denominator  int
	// 	sum          int
	// 	NOK          int
	// 	a            int
	// 	b            int
	// 	m            int
	// 	denominatorA int
	// 	numeratorA   int
	// 	c int
	// 	d int
	// 	NOD int
	// )
	// fmt.Scan(&count)
	// for i := 0; i < count; i++ {
	// 	fmt.Scan(&numerator, &denominator)
	// 	if i == 0 {
	// 		numeratorA = numerator
	// 		denominatorA = denominator
	// 		NOK = denominator
	// 		sum = numerator
	// 	} else {
	// 		a = NOK
	// 		b = denominator
	// 		m = a * b
	// 		for a != 0 && b != 0 {
	// 			if math.Abs(float64(a)) > math.Abs(float64(b)) {
	// 				a = a % b
	// 			} else {
	// 				b = b % a
	// 			}
	// 		}
	// 		NOK = m / (a + b)
	// 		sum = NOK / denominator * numerator + NOK / denominatorA * numeratorA
	// 		numeratorA = sum
	// 		denominatorA = NOK
	// 	}
	// }
	// c = sum
	// d = NOK
	// for c != 0 && d != 0 {
	// 	if math.Abs(float64(c)) > math.Abs(float64(d)) {
	// 		c = c % d
	// 	} else {
	// 		d = d % c
	// 	}
	// }
	// NOD = c + d
	// sum = sum / NOD
	// NOK = NOK / NOD
	// fmt.Printf("%d/%d", sum, NOK)

	// fmt.Scan(&a, &b)
	// m = a * b
	// for a != 0 && b != 0 {
	// 	if math.Abs(float64(a)) > math.Abs(float64(b)) {
	// 		a = a % b
	// 	} else {
	// 		b = b % a
	// 	}
	// }
	// fmt.Printf("НОД: %d\n", a+b)
	// fmt.Printf("НОК: %d", m / (a + b))
	// проверить вводимое значение ноль

	var (
		value int
		a     bool
	)

	fmt.Scan(&value)

	for i := 1; i <= value; i++ {

		switch {
		case i == 1 || i == value:
			fmt.Printf("%d ", i)
		case i > 1:
			if value%i == 0 {
				a = true
				fmt.Printf("%d ", i)
			}
		}
		

	}
	if a == false && value != 1 {
		fmt.Printf ("\nACHTUNG")
	}
}
