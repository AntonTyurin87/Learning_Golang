// Перебираем все позиции и для каждой позиции в строке ещё раз перебираем все позиции и в случае совпадения прибавим к счетчику единицу. Найдём максимальное значение счетчика.
// O(N*N)
package main

import "fmt"

func main() {
	s := "fgfjhinscuueddhh"
	var ans string //Переменная для самого частого символа
	var anscnt int //Переменная для количества символов

	//Перебираем каждый символ и ищем совпадение с ним каждый раз
	for i, _ := range s {
		nowcnt := 0
		for j, _ := range s {
			if s[i] == s[j] {
				nowcnt++
			}
			if nowcnt > anscnt {
				ans = string(s[j])
				anscnt = nowcnt
			}
		}
	}

	//Выводим результат
	fmt.Println(ans)
}
