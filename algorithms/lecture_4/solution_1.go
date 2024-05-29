//Дано два числа X и Y без ведущих нулей
//Необходимо проверить, можно ли получить первое из второрго перестановкой цифр

//Решение
//Посчитаем количество вхождений каждой цифры в каждое из чисел и сравним. Цифры будем постепенно добывать из числа справа с помощью %10 и //10

package main

import "fmt"

func isDigitPermutation(x, y int) bool {
	digitsX := numsMap(x)
	digitsY := numsMap(y)

	for i := 0; i <= 9; i++ {
		if digitsX[i] != digitsY[i] {
			return false
		}
	}
	return true
}

func main() {
	x := 2021
	y := 1202

	fmt.Println(isDigitPermutation(x, y))
}

func numsMap(x int) map[int]int {
	numMap := make(map[int]int)

	for i := 0; i <= 9; i++ {
		numMap[i] = 0
	}

	for x > 0 {
		lastdigit := x % 10
		numMap[lastdigit]++
		x /= 10
	}
	return numMap
}
