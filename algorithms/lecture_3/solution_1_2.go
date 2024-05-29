//Дана последовательность положительных чисел длиной N и число X
//Нужно найти два различных числа A и B из последовательности, такие что A + B = X
//или вернуть пару 0, 0, если такой пары чисел нет

//Решение за O(N)
//Будем хранить все уже обработанные числа в множестве.
//Если очередное числло nownum, а X - nownum есть в множестве, то мы нашли слагаемое.

package main

import "fmt"

// Реализация через map
func twoTermSwithSumX_2(nums []int, x int) (int, int) {
	prevnum := make(map[int]bool)
	for _, nownum := range nums {
		if _, ok := prevnum[x-nownum]; ok {
			return x - nownum, nownum
		} else {
			prevnum[nownum] = true
		}
	}
	return 0, 0
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	x := 4

	fmt.Println(twoTermSwithSumX_2(nums, x))
}
