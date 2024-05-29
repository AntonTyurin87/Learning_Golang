//На шахматной доске NxN находятся M ладей (ладья бьёт клетки на той же горизонтали или вертикали до ближайшей занятой).
//Определите, сколько пар ладей бьют друг друга. Ладьи задаются парой чисел I и J, обозначающих координаты клетки.
//1<= N <= (10 в 9), 0<= M <= 2x(10 в 5)

//Решение O(M)
//Для каждой занятой горизонтали и вертикали будем хранить количество ладей на них.
//Количество пар в горизонтали(вертикали) равно количеству ладей минус 1.
//Суммируем это количество пар для всех горизонталей и вертикалей.

package main

import "fmt"

func countBeatingRooks(rookcoords [][]int) int {
	rooksInRow := make(map[int]int)
	rooksInCol := make(map[int]int)

	for _, coord := range rookcoords {
		rooksInRow = addRoock(rooksInRow, coord[0])
		rooksInCol = addRoock(rooksInCol, coord[1])
	}
	return countpairs(rooksInRow) + countpairs(rooksInCol)
}

func main() {
	rookcoords := [][]int{[]int{1, 2}, []int{5, 8}, []int{13, 98}, []int{1, 96}}

	fmt.Println(countBeatingRooks(rookcoords))

}

func addRoock(roworcol map[int]int, key int) map[int]int {
	if _, ok := roworcol[key]; !ok {
		roworcol[key] = 1
	} else {
		roworcol[key]++
	}
	return roworcol
}

func countpairs(roworcol map[int]int) int {
	pairs := 0
	for _, values := range roworcol {
		pairs += values - 1
	}
	return pairs
}
