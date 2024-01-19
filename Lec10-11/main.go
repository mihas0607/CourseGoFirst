package main

import (
	"fmt"
)

func main() {
	// var val string
	// input := bufio.NewScanner(os.Stdin)
	// for {
	// 	if input.Scan() {
	// 		val = input.Text()
	// 		if val == "" {
	// 			fmt.Println("НЕ СОГЛАСЕН")
	// 			break
	// 		}
	// 		runeSlice := []rune(val)
	// 		if ((runeSlice[0] == 'д' || runeSlice[0] == 'Д') && (runeSlice[len(runeSlice)-1] == 'а' || runeSlice[len(runeSlice)-1] == 'А')) || (runeSlice[0] == 'А' || runeSlice[0] == 'а') && (runeSlice[len(runeSlice)-1] == 'д' || runeSlice[len(runeSlice)-1] == 'Д') {
	// 			fmt.Println("СОГЛАСЕН")
	// 			break
	// 		} else {
	// 			fmt.Println("НЕ СОГЛАСЕН")
	// 			break
	// 		}
	// 	}
	// }

	// var (
	// 	word      string
	// 	startWord string
	// )
	// input := bufio.NewScanner(os.Stdin)
	// input.Scan()
	// startWord = input.Text()
	// for {
	// 	if input.Scan() {
	// 		word = input.Text()
	// 		runeStartWord := []rune(startWord)
	// 		runeWord := []rune(word)
	// 		if runeWord[0] == runeStartWord[len(runeStartWord)-1] {
	// 			startWord = word
	// 		} else if runeStartWord[len(runeStartWord)-1] == 'ь' && runeWord[0] == runeStartWord[len(runeStartWord)-2] {
	// 			startWord = word
	// 		} 	else {
	// 			fmt.Println(word)
	// 			break
	// 		}
	// 	}
	// }

	// var (
	// 	result    string
	// 	newResult string
	// )
	// input := bufio.NewScanner(os.Stdin)
	// input.Scan()
	// result = input.Text()
	// runeResult := []rune(result)
	// runeNewResult := []rune(newResult)
	// count := 0
	// fmt.Println(string(runeResult))
	// for len(runeResult) > 2 {
	// 	if count%2 != 0 {
	// 		for i, v := range runeResult {
	// 			if i < len(runeResult)-2 {
	// 				runeNewResult = append(runeNewResult, v)
	// 			}
	// 		}
	// 		runeResult = runeNewResult
	// 		fmt.Println(string(runeResult))
	// 		runeNewResult = []rune("")
	// 		count++
	// 	} else {
	// 		for i, v := range runeResult {
	// 			if i > 1 {
	// 				runeNewResult = append(runeNewResult, v)
	// 			}
	// 		}
	// 		runeResult = runeNewResult
	// 		fmt.Println(string(runeResult))
	// 		runeNewResult = []rune("")
	// 		count++
	// 	}
	// }

	var (
		val   string
		count int
		a     bool
		b     int
	)
	fmt.Scan(&val)
	runeSlice := []rune(val)
	for _, v := range runeSlice {
		if v == 'о' {
			if a {
				count++
			}
			if b <= count {
				b = count + 1
			}
			a = true
		} else {
			count = 0
			a = false
		}
	}
	fmt.Println(b)

	mihasMap := make(map[string]int)
	mihasMap["Mihas"] = 37
	mihasMap["aKSANA"] = 33
	fmt.Println(mihasMap)

	aksanaMap := map[string]int{
		"Passha": 37,
		"Lena":   30,
	}
	fmt.Println(aksanaMap)
	if value, ok := mihasMap["Den"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Does not exist")
	}

	fmt.Println(len(aksanaMap))

}
