//Сортировка подсчетом
//Последовательность непустая
//Сортируем копию

package main

import "fmt"

// решение через map
func countSort(seq []int) []int {
	minVal := min(seq)
	maxVal := max(seq)
	counter := make(map[int]int)

	var result []int

	for i := minVal; i <= maxVal; i++ {
		counter[i] = 0
	}

	for _, j := range seq {
		counter[j]++
	}

	for i := minVal; i <= maxVal; i++ {
		for k := 0; k < counter[i]; k++ {
			result = append(result, i)
		}
	}
	return result
}

// На примере школьных оценок
func main() {
	seq := []int{1, 2, 3, 3, 3, 4, 4, 5, 5, 5, 5, 2, 4}

	fmt.Println(countSort(seq))
}

// минимальное число в последовательности
func min(seq []int) int {
	min := seq[0]
	for i := 1; i < len(seq); i++ {
		if min > seq[i] {
			min = seq[i]
		}
	}
	return min
}

// максимальное число в последовательности
func max(seq []int) int {
	max := seq[0]
	for i := 1; i < len(seq); i++ {
		if max < seq[i] {
			max = seq[i]
		}
	}
	return max
}
