//Дана последовательность чисел длиной N
//Найти минимальное четное число в последовательности или вывести -1, если такого числа не существует.

package main

import "fmt"

func findMinEven(seq []int) int {
	ans := -1
	flag := false //с переменной flag работает и на четные и на нечетные значения.

	for _, i := range seq {
		if i%2 == 0 && (!flag || i < ans) {
			ans = i
			flag = true
		}
	}
	return ans
}

func main() {

	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10}

	fmt.Println(findMinEven(seq))
}
