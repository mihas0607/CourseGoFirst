package main

import (
	"fmt"
)

func main() {
	// floatArr := [...]float64{12, 11.3, 115.6}
	// for i := 0; i < len(floatArr); i++ {
	// 	fmt.Printf("%d element of arr is %.2f\n", i, floatArr[i])
	// }
	// var sum float64
	// for id, val := range floatArr {
	// 	fmt.Printf("%d element of arr is %.2f\n", id, val)
	// 	sum += val
	// }
	// fmt.Println(sum)
	// for _, val := range floatArr {
	// 	fmt.Printf("element of arr is %.2f\n", val)
	// 	sum += val
	// }
	// words := [2][2]string{
	// 	{"Bob", "Alice"},
	// 	{"Mihas", "Petr"},
	// }
	// for _, val1 := range words {
	// 	for _, val2 := range val1 {
	// 		fmt.Printf("%s ", val2)
	// 	}
	// 	fmt.Println()
	// }

	// var (
	// 	kekArr []string
	// 	myArr  []string
	// 	c      bool
	// )
	// a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// a = strings.TrimRight(a, "\r\n")
	// a1, _ := strconv.Atoi(a)
	// b, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// b = strings.TrimRight(b, "\r\n")
	// b1, _ := strconv.Atoi(b)
	// for i := 0; i < a1; i++ {
	// 	val1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// 	val1 = strings.TrimRight(val1, "\r\n")
	// 	kekArr = append(kekArr, val1)
	// }
	// for i := 0; i < b1; i++ {
	// 	val2, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// 	val2 = strings.TrimRight(val2, "\r\n")
	// 	myArr = append(myArr, val2)
	// }
	// for _, val2 := range myArr {
	// loop2:
	// 	for _, val1 := range kekArr {
	// 		switch {
	// 		case val2 == val1:
	// 			c = true
	// 			break loop2
	// 		default:
	// 			c = false
	// 		}
	// 	}
	// 	if c {
	// 		fmt.Println("ЕСТЬ")
	// 	} else {
	// 		fmt.Println("НЕТ В НАЛИЧИИ")
	// 	}
	// }

	// var (
	// 	numberOfValues int
	// 	val            int
	// 	a int
	// 	b int
	// 	sum int
	// 	slice []int
	// )
	// fmt.Scan(&numberOfValues)
	// for i := 0; i < numberOfValues; i++ {
	// 	fmt.Scan(&val)
	// 	slice = append(slice, val)
	// }
	// fmt.Scan(&a,&b)
	// for i, v := range slice {
	// 	if i >= a-1 && i <= b-1 {
	// 		sum += v
	// 	}
	// }
	// fmt.Println(sum)

	// var (
	// 	numberOfTasks     int
	// 	slice             []string
	// 	numberOfTasksToDo int
	// 	sliceToDo         []int
	// )
	// fmt.Scan(&numberOfTasks)
	// for i := 0; i <= numberOfTasks; i++ {
	// 	val, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// 	val = strings.TrimRight(val, "\r\n")
	// 	slice = append(slice, val)
	// }
	// fmt.Scan(&numberOfTasksToDo)
	// for i := 0; i <= numberOfTasksToDo; i++ {
	// 	a, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// 	a = strings.TrimRight(a, "\r\n")
	// 	a1, _ := strconv.Atoi(a)
	// 	sliceToDo = append(sliceToDo, a1)
	// }
	// for _, v1 := range sliceToDo {
	// 	for i2, v2 := range slice {
	// 		switch {
	// 		case v1 == i2 && v1 > 0:
	// 			fmt.Println(v2)

	// 		}
	// 	}
	// }

	// var (
	// 	a int
	// 	b string
	// )
	// fmt.Scan(&a)
	// arr := [...]string{
	// 	for i := 0; i < a; i++ {
	// 		fmt.Scan(&b)
	// 		arr[i] = b
	// 	}
	// }
	
	// fmt.Println(arr)

	// originArr := [...]int{30, 40, 50, 60, 70, 80}
	// firstSlice := originArr[1:4] // Это набор ссылок на элементы нижележащего массива
	// for i, _ := range firstSlice {
	// 	firstSlice[i]++
	// }
	// fmt.Println("OriginArr:", originArr)
	// fmt.Println("FirstSlice:", firstSlice)

	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // В этот момент numSlice больше не ссылается на numArr
	numSlice[0] = 10
	fmt.Println(numArr)
	fmt.Println(numSlice)
}
