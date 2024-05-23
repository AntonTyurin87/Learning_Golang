//Дана последовательность чисел длиной N (N>0)
//Найти максимальное число в последовательности

package main

import "fmt"

func findx_3(seq []int) int {
	//Значения могут быть большими, и, чтобы не копирывать их при каждом сравнении передавать будем указатель на них.
	var ans *int

	ans = &seq[0]

	for _, v := range seq[1:] {
		if v > *ans {
			ans = &v
		}
	}
	return *ans
}

func main() {
	seq := []int{1, 2, 1, 3, 3}

	fmt.Println(findx_3(seq))

}
