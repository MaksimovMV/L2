package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type arguments struct {
	after      *int
	before     *int
	context    *int
	count      *bool
	ignoreCase *bool
	invert     *bool
	fixed      *bool
	lineNum    *bool
}

// grep - выводит найденные и необходимые связанные строки в stdout
func grep(arr []string, pattern string, args arguments) error {

	nums, err := findNumsOfLines(arr, pattern, args)
	if err != nil {
		return err
	}
	if *args.count {
		fmt.Println(len(nums))
		return nil
	}
	for _, i := range nums {
		if *args.lineNum {
			fmt.Printf("%v ", i+1)
		}
		fmt.Print(arr[i])
	}
	return nil
}

// findNumsOfLines - сохраняет номера необходимых строк в срез
func findNumsOfLines(arr []string, pattern string, args arguments) (numOfLines []int, err error) {
	linesBefore := 0
	linesAfter := 0
	if *args.fixed {
		pattern = regexp.QuoteMeta(pattern)
	}
	if *args.ignoreCase {
		pattern = "(?i)" + pattern
	}
	if *args.before > *args.context {
		linesBefore = *args.before
	} else {
		linesBefore = *args.context
	}
	if *args.after > *args.context {
		linesAfter = *args.after
	} else {
		linesAfter = *args.context
	}

	for i, s := range arr {
		match, err := regexp.MatchString(pattern, s)
		if err != nil {
			return nil, err
		}
		if *args.invert {
			match = !match
		}
		if match {
			for j := i - linesBefore; j <= i+linesAfter; j++ {
				if j < 0 || (len(numOfLines) != 0 && j <= numOfLines[len(numOfLines)-1]) {
					continue
				} else if j > len(arr)-1 {
					break
				}
				numOfLines = append(numOfLines, j)
			}
		}
	}

	return numOfLines, nil
}

func main() {
	args := arguments{
		after:      flag.Int("A", 0, "печатать +N строк после совпадения"),
		before:     flag.Int("B", 0, "печатать +N строк до совпадения"),
		context:    flag.Int("C", 0, "(A+B) печатать ±N строк вокруг совпадения"),
		count:      flag.Bool("c", false, "количество строк"),
		ignoreCase: flag.Bool("i", false, "игнорировать регистр"),
		invert:     flag.Bool("v", false, "вместо совпадения, исключать"),
		fixed:      flag.Bool("F", false, "точное совпадение со строкой, не паттерн"),
		lineNum:    flag.Bool("n", false, "печатать номер строки"),
	}
	//var args arguments
	//
	//args.lineNum = flag.Bool("n", false, "печатать номер строки")

	flag.Parse()

	arguments := flag.Args()
	if len(arguments) < 1 {
		fmt.Println("Need file name!")
		return
	}
	fileName := arguments[len(arguments)-1]
	pattern := arguments[len(arguments)-2]

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	arr := strings.SplitAfter(string(all), "\n")

	if err := grep(arr, pattern, args); err != nil {
		panic(err)
	}
}
