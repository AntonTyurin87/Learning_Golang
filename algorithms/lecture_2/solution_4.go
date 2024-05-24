//Дана последовательность чисел длиной N (N>1)
//Найти максимальное число в последовательности и второе по величине число
//(такое, которое будет максимальным, если вычеркнуть из последовательности, если вычеркнуть из последовательности одно максимальное число).

package main

import (
	"fmt"
	"math"
)

func findMax2(seq []int) (int, int) {

	max1 := int(math.Max(float64(seq[0]), float64(seq[1])))
	max2 := int(math.Min(float64(seq[0]), float64(seq[1])))

	for _, i := range seq[2:] {
		if i > max1 {
			max2 = max1
			max1 = i
		} else if seq[i] > max2 {
			max2 = i
		}
	}
	return max1, max2
}

func main() {

	seq := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 10}

	fmt.Println(findMax2(seq))

}
