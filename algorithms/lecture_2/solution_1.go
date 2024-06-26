// Дана последовательность чисел длиной N.
// Найти первое (левое) вхождение положительного числа X в неё или вывести -1, если число X не встречалось.
package main

import "fmt"

func findx_1(seq []int, x int) int {
	ans := -1

	for i, v := range seq {
		if v == x {
			ans = i
			return ans
		}
	}

	return ans
}

func main() {

	x := 2
	seq := []int{1, 2, 1, 3, 3}

	fmt.Println(findx_1(seq, x))
}
