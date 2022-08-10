package main

import (
	"fmt"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// toNum - переводит руну в int
func toNum(r rune) (int, error) {
	if n, err := strconv.Atoi(string(r)); err != nil {
		return 0, err
	} else {
		return n, nil
	}
}

// unpackString - распаковка строк
func unpackString(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	if _, err := toNum(runes[0]); err == nil {
		return ""
	}
	var result []rune

	for i := 0; i < len(runes)-1; i++ {
		if n, err := toNum(runes[i+1]); err != nil {
			result = append(result, runes[i])
		} else {
			t := i
			for j := i + 2; j < len(runes); j++ {
				if k, err := toNum(runes[j]); err != nil {
					break
				} else {
					n = n*10 + k
					i++
				}
			}
			for l := 0; l < n; l++ {
				result = append(result, runes[t])
			}
			i++
		}
	}
	if _, err := toNum(runes[len(runes)-1]); err != nil {
		result = append(result, runes[len(runes)-1])
	}

	return string(result)
}

func main() {
	fmt.Println(unpackString(""))
}
