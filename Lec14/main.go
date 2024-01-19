package main

import (
	"fmt"
)

var (
	n int
	m int
)

// func factorial(n uint64) uint64 {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }

func factorial(n int) uint64 {
	var fact uint64 = 1
	for i := 1; i <= n; i++ {
		fact *= uint64(i)
	}
	// fmt.Print(fact)
	return fact
}

func combination(n int, m int) int {
	facN := factorial(n)
	facM := factorial(m)
	facNminM := factorial(n - m)
	facNfacNminM := facM * facNminM
	fmt.Println(facN, facM, facNminM, facNfacNminM)
	return int(facN / facNfacNminM)
	// return int(factorial(n) / (factorial(m) * factorial(n-m)))
}

// func combination(n int, m int) int {
// 	facN := 1
// 	facM := 1
// 	facNminusM := 1
// 	for i := 1; i <= n; i++ {
// 		facN *= i
// 	}
// 	for i := 1; i <= m; i++ {
// 		facM *= i
// 	}
// 	for i := 1; i <= n-m; i++ {
// 		facNminusM *= i
// 	}
// 	return facN / (facM * facNminusM)
// }

func main() {
	fmt.Scan(&n, &m)
	if n <= 0 {
		return
	}
	if m <= 0 {
		return
	}
	if n < m {
		return
	}
	// fmt.Println(combination(n, m))
	fmt.Println(factorial(n))
	// fmt.Println(combination(n, m))
	fmt.Println(factorial(m))
}
