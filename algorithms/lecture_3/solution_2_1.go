//Дан словарь из N слов, длина каждого не превосходит K
//В записи каждого из M слов текста (каждое длиной K) может быть пропущена одна буква.
//Для каждого слова сказать, входит ли оно (возможно с одной пропущенной буквой), в словарь.

//Решение за O(NK*K + M)
//Выбросим из каждого слова словаря по одной букве всеми возможными способами за O(NK) и положим получившиеся слова в множества.
//Для каждого слова из текста просто проверим, есть ли оно в словаре за O(1)

package main

import (
	"fmt"
	"strings"
)

// Реализация через map, а текстуже разобран на слова
func wordSinDict(dictionary []string, text []string) []bool {
	goodwords := make(map[string]bool)
	var ans []bool

	//Вносим в map первичный словарь
	for _, v := range dictionary {
		goodwords[v] = true
	}

	//Вносим в map слова без одной буквы
	for _, word := range dictionary {
		for delpos, _ := range word {
			newWord := strings.Join([]string{string(word[:delpos]), string(word[delpos+1:])}, "")
			goodwords[newWord] = true
		}
	}

	//Сверяем наличие слов в словаре
	for _, word := range text {
		if _, ok := goodwords[word]; ok {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans
}

func main() {

	dictionary := []string{"algoritm", "road", "walked", "Sasha", "along", "and", "chew", "dryer"} //Слова "highway" заведомо нет

	text := []string{"walked", "Ssha", "along", "highway", "and", "chew", "dryer"} //Допущены ошибки в словах "Sasha" и "chew"

	fmt.Println(wordSinDict(dictionary, text))
}
