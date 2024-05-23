// Перебем все символы, встречающиеся в строке, а затем переберём все позиции и в случае совпадения прибавим к счетчику единицу. Найдём максимальное значение счетчика.
// O(N*k)
// Адоптация через map
package main

import "fmt"

func main() {
	s := "fgfjhinscuueddhh"

	var ans string //Переменная для самого частого символа
	var anscnt int //Переменная для количества символов

	uniq := set(s) //Слпйс уникальных рун

	//Считаем количество уникальных значений
	for _, i := range uniq {
		nowcnt := 0
		for _, j := range s {
			if i == j {
				nowcnt++
			}
			if nowcnt > anscnt {
				ans = string(j)
				anscnt = nowcnt
			}

		}
	}

	//Выводим результат
	fmt.Println(ans)
}

// Ревлизация функции, которая возвращает уникальные значения после первого прохода по строке
func set(str string) []rune {

	var strSlice []rune //Возвращаемая строка

	set := make(map[rune]bool)

	for _, v := range str {
		if _, ok := set[v]; !ok {
			set[v] = true
		}

		for k, _ := range set {
			strSlice = append(strSlice, k)
		}

	}
	return strSlice
}
