package main

import (
	"fmt"
	"sort"
	"unicode/utf8"
)

var (
	countOfStrings int
	myString       string
)

type SmartList struct {
	words []string
}

func New(newWords []string) *SmartList {
	return &SmartList{newWords}
}

// Sort words and printed it
func (sl *SmartList) GetAnswer() {
	for _, v := range sl.words {
		fmt.Println(v)
	}
}

func (sl *SmartList) DefaultSort() *SmartList{
	sort.Strings(sl.words)
	return &SmartList{sl.words}
}

func (sl *SmartList) BubleSortNew() *SmartList {
	for i := 0; i < len(sl.words)-1; i++ {
		for j := i + 1; j < len(sl.words); j++ {
			if utf8.RuneCountInString(sl.words[i]) > utf8.RuneCountInString(sl.words[j]) {
				sl.words[i], sl.words[j] = sl.words[j], sl.words[i]
			}
		}
	}
	return &SmartList{sl.words}
}

func main() {
	newWords := make([]string, 0, 15)
	fmt.Scan(&countOfStrings)
	for i := 0; i < countOfStrings; i++ {
		fmt.Scan(&myString)
		newWords = append(newWords, myString)
	}
	New(newWords).DefaultSort().BubleSortNew().GetAnswer()
}

// func BubleSort(slice []string) []string {
// 	for i := 0; i < len(slice)-1; i++ {
// 		for j := i + 1; j < len(slice); j++ {
// 			if utf8.RuneCountInString(slice[i]) > utf8.RuneCountInString(slice[j]) {
// 				slice[i], slice[j] = slice[j], slice[i]
// 			}
// 		}
// 	}
// 	return slice
// }
