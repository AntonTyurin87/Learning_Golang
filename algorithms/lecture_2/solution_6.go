//Дана последовательность слов
//Вывести все самые короткие слова через пробел

package main

import (
	"fmt"
	"strings"
)

func shortWords(words []string) string {

	minLen := len(words[0])

	var answerSlice []string

	for _, i := range words {
		if len(i) < minLen {
			minLen = len(i)
		}
	}

	for _, i := range words {
		if len(i) == minLen {
			answerSlice = append(answerSlice, i)
		}
	}
	return strings.Join(answerSlice, " ")
}

func main() {

	words := []string{"qqq", "wwww", "ee", "gg", "ddeff", "rr"}

	fmt.Println(shortWords(words))

}
