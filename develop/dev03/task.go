package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sort(k int, n, r, u bool, arr [][]string) ([][]string, error) {
	f := func(x, y string) bool { return x < y }

	if k > len(arr[0]) || k <= 0 {
		return nil, fmt.Errorf("invalid field number")
	}
	if n && r {
		f = byNumRev
	} else if n {
		f = byNum
	} else if r {
		f = func(x, y string) bool { return x > y }
	}
	result := sortStringsByField(arr, f, k-1)
	if u {
		result = removeDuplicateStr(result)
	}

	return result, nil
}

// sortStringsByField - сортирует строку по полям по правилу в функции f
func sortStringsByField(a [][]string, f func(x, y string) bool, n int) [][]string {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivotIndex := rand.Int() % len(a)

	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	for i := range a {
		if f(a[i][n], a[right][n]) {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	sortStringsByField(a[:left], f, n)
	sortStringsByField(a[left+1:], f, n)

	return a
}

// removeDuplicateStr - удаляет дупликаты в отсортированном срезе
func removeDuplicateStr(sortedSlice [][]string) [][]string {
	if len(sortedSlice) <= 1 {
		return sortedSlice
	}
	for i := 1; i < len(sortedSlice); i++ {
		if reflect.DeepEqual(sortedSlice[i], sortedSlice[i-1]) {
			sortedSlice = append(sortedSlice[:i], sortedSlice[i+1:]...)
			i--
		}
	}

	return sortedSlice
}

// toNum - переводит руну в int
func toNum(r rune) (int, bool) {
	if n, err := strconv.Atoi(string(r)); err != nil {
		return 0, false
	} else {
		return n, true
	}
}

// takeNum - извлекает число из начала строки
func takeNum(input string) (n int, s string) {
	s = input
	n = 0
	for i := 0; i < len(input); i++ {
		k, ok := toNum(rune(input[i]))
		if ok {
			n = n*10 + k
			if i != len(input)-1 {
				s = input[i+1:]
			} else {
				s = ""
			}
		} else {
			break
		}
	}
	return n, s
}

// byNum - определяет правило для сортировки
func byNum(x, y string) bool {
	numX, strX := takeNum(x)
	numY, strY := takeNum(y)
	if numX == numY {
		return strX < strY
	} else {
		return numX < numY
	}
}

// byNumRev - определяет правило для сортировки
func byNumRev(x, y string) bool {
	numX, strX := takeNum(x)
	numY, strY := takeNum(y)
	if numX == numY {
		return strX > strY
	} else {
		return numX > numY
	}
}

// writeArrToFile - запись среза в файл построчно
func writeArrToFile(result [][]string) error {
	var b bytes.Buffer

	for i := 0; i < len(result); i++ {
		s := strings.Join(result[i], " ")
		if _, err := b.WriteString(s + "\n"); err != nil {
			return err
		}
		fmt.Println(s)
	}

	if err := ioutil.WriteFile("result.txt", b.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

// readArrFromFile - чтение файла, разбиение содержимого на поля и запись в срез
func readArrFromFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file due to error: %v\n", err)
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file due to error: %v\n", err)
	}
	arr1 := strings.SplitAfter(string(all), "\n")
	arr2 := make([][]string, 0)
	for _, a := range arr1 {
		arr2 = append(arr2, strings.Fields(a))
	}
	return arr2, nil
}

func main() {
	r := flag.Bool("r", false, "reverse the result of comparisons")
	k := flag.Int("k", 1, "number of column")
	u := flag.Bool("u", false, "need to remove duplicates")
	n := flag.Bool("n", false, "compare according to string numerical value")

	flag.Parse()

	arguments := flag.Args()
	if len(arguments) < 1 {
		fmt.Println("Need file name!")
		return
	}
	fileName := arguments[len(arguments)-1]

	arr, err := readArrFromFile(fileName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to read array from file due to error: %v\n", err)
		os.Exit(1)
	}

	result, err := sort(*k, *n, *r, *u, arr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed sort array due to error: %v\n", err)
		os.Exit(1)
	}

	if err := writeArrToFile(result); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write result array to file due to error: %v\n", err)
		os.Exit(1)
	}
}
