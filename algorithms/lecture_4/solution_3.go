//Выведите гистогрмму как в примере (коды символов отсортированы)
//S = Hello, word!
//
//      #
//      ##
//##########
// !,Hdelorw

//Решение
//Для каждого символа в словаре подсчитаем, сколько раз он встречается.
//Найдем самый частый символ и переберём количество от этого числа до 1.
//Пройдём по всем отсортированным ключам и если количество больше счетчика - выведем #

package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func printChart(s string) {
	symCount, maxSymCount := symCountAndMaxSymCount(s)

	sorteDuniqsyms := sortKeys(symCount)

	for row := maxSymCount; row >= 0; row-- {
		for _, sym := range sorteDuniqsyms {
			if symCount[sym] >= row {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println(strings.Join(sorteDuniqsyms, ""))
}

func main() {
	s := "Hello, world!"
	printChart(s)
}

func symCountAndMaxSymCount(s string) (map[string]int, int) {
	symCount := make(map[string]int)
	maxCount := 0
	for _, i := range s {
		if _, ok := symCount[string(i)]; !ok {
			symCount[string(i)] = 0

		} else {
			symCount[string(i)]++
			maxCount = int(math.Max(float64(symCount[string(i)]), float64(maxCount)))
		}
	}
	return symCount, maxCount
}

func sortKeys(inMap map[string]int) []string {
	var inString []string
	for key, _ := range inMap {
		inString = append(inString, key)
	}
	sort.Strings(inString)
	return inString
}
