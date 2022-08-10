package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

//findAnagram - поиск анаграмм
func findAnagram(arr *[]string) *map[string][]string {
	anagrams := make(map[string][]string) //ключ - отсортированная строка по символам
iterating:
	for _, word := range *arr {
		word = strings.ToLower(word)
		wordRunes := []rune(word)
		sort.Slice(wordRunes, func(i, j int) bool {
			return wordRunes[i] < wordRunes[j]
		})
		if _, ok := anagrams[string(wordRunes)]; !ok {
			anagrams[string(wordRunes)] = []string{word}
			continue
		}
		for _, w := range anagrams[string(wordRunes)] {
			if w == word {
				continue iterating
			}
		}
		anagrams[string(wordRunes)] = append(anagrams[string(wordRunes)], word)
	}
	result := make(map[string][]string) //ключ - первое встретившееся в словаре слово из множества, срез значений отсортирован

	for _, value := range anagrams {
		if len(value) < 2 {
			continue
		}
		key := value[0]
		sort.Strings(value)
		result[key] = value
	}
	return &result
}

func main() {
	arr := []string{"пятка", "пятак", "тяпка", "листок", "слиток", "пятка", "столик"}
	fmt.Println(findAnagram(&arr))
}
