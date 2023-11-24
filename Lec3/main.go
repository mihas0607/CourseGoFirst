package main

import (
	"fmt"
	// "math"
)

func main() {
	var (
	// a     bool
	// b     float64
	// perimeter int
	// area      int
	)

	// fmt.Scan(&name)
	// fmt.Scan(&age)
	// fmt.Scan(&name, &age)

	// fmt.Printf("My name is: %s\nMy age is: %d\n", name, age)

	// fmt.Fscan(os.Stdin, &age)
	// fmt.Println("New age: ", age)

	// fmt.Scan(&drink, &meal, &subMeal, &time)
	// fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'", drink, meal, subMeal, time)

	// fmt.Scan(&a, &b)
	// fmt.Println(math.Pow(a+b, 2))

	// fmt.Println(a)

	// var uid = "MIHAS"
	// fmt.Println("My uid:", uid)
	// fmt.Printf("%T\n", uid)

	// aArg, bArg := 10, 30
	// fmt.Println(aArg, bArg)
	// aArg, bArg = 30, 40
	// fmt.Println(aArg, bArg)

	// var (
	// 	personName string = "Mihas"
	// 	personAge  int    = 37
	// 	personUid  int
	// )

	

	// fmt.Printf("Name: %s\nAge: %d\nUID: %d\n", personName, personAge, personUid)

	var (
		age  int
		name string
	)

	// fmt.Scan(&age)
	// fmt.Scan(&name)
	fmt.Scan(&age, &name)

	fmt.Printf("My name is: %s\nMy age is : %d\n", name, age)

}
