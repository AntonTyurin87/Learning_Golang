//Дана последовательность положительных чисел длиной N и число X
//Нужно найти два различных числа A и B из последовательности, такие что A + B = X
//или вернуть пару 0, 0, если такой пары чисел нет

//Решение за O(N*N)
//Переберём числа A за O(N). Переберём числа B за O(N).
//Если их сумма равна X, то вернём эту пару.

package main

import "fmt"

func twoTermSwithSumX_1(nums []int, x int) (int, int) {
	for ai, _ := range nums {
		for bi, _ := range nums[ai:] {
			if nums[ai]+nums[bi] == x {
				return nums[ai], nums[bi]
			}
		}
	}
	return 0, 0
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	x := 3

	fmt.Println(twoTermSwithSumX_1(nums, x))
}
