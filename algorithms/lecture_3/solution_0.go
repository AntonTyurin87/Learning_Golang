//Наше собственное (МУЛЬТИ)множество

package main

import "fmt"

// Добавление элемента массива
func add(x int, array [10][]int, setsize int) [10][]int {
	array[x%setsize] = append(array[x%setsize], x)
	return array
}

// Поиск элемента массива
func find(x int, array [10][]int, setsize int) bool {
	for _, i := range array[x%setsize] {
		if i == x {
			return true
		}
	}
	return false
}

// Удаление элемента массива
func delete(x int, array [10][]int, setsize int) [10][]int {
	xlist := array[x%setsize]
	for i, _ := range xlist {
		if xlist[i] == x {
			xlist[i], xlist[len(xlist)-1] = xlist[len(xlist)-1], xlist[i]
			xlist[len(xlist)-1] = 0 //Записываем в слайс на последнюю позицию 0. Это работает, если у нас не будет нулей в массиве((
		}
	}
	return array
}

func main() {
	setsize := 10
	var myset [10][]int //Здесь длинна массива должна быть равнв значению setsize

	myset = add(3, myset, setsize)
	myset = add(4, myset, setsize)

	fmt.Println(find(5, myset, setsize))
	fmt.Println(find(3, myset, setsize))

	myset = delete(4, myset, setsize)
	myset = delete(3, myset, setsize)

	fmt.Println(find(3, myset, setsize))
}
