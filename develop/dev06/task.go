package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func cut(a string, f int, d string, s bool) (string, error) {
	split := strings.Split(a, d)
	if s {
		if len(split) == 1 {
			return "", nil
		}
	}
	if f > len(split) || f < 1 {
		return "", fmt.Errorf("некорректное поле f")
	}
	return split[f-1], nil
}

func main() {
	fields := flag.Int("f", 1, "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() && s.Text() != "стоп" {
		a := s.Text()
		l, err := cut(a, *fields, *delimiter, *separated)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(l)
	}
}
