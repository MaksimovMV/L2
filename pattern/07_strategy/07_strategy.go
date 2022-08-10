package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

//Стратегия - это поведенческий паттерн проектирования, который позволяет менять логику работы в зависимости от выбранной стратегии (на стороне клиента)
// Применение - 1) нужно использовать разные вариации какого-то алгоритма вниутри одного объекта
// 2) есть множество похожих классов, отличающихся только некоторым поведением
// 3) не хотите обнажать детали реализации для других классов
// 4) когда различные вариации алгоритмов реализованы в виде развесистого
// условного оператора. Каждая ветка такого оператора преставляет собой вариацию алгоритма
// Плюсы - 1) горячая замена алгоритмов на лету
// 2) изолирует код и данные алгоритмов от остальных классов
// 3) уход от наследования к делегированию
// 4) принцип открытости/закрытости
// минусы - 1) усложняет программу за счет доп классов
// 2) клиент должен знать, в чем состоит разница между стратегиями, чтобы выбрать подходящую

// StrategySort - интерфейс стратегии
type StrategySort interface {
	Sort([]int)
}

// BubbleSort - конкретная стратегия
type BubbleSort struct {
}

func (s *BubbleSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}

	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// InsertionSort - конкретная стратегия
type InsertionSort struct {
}

func (s *InsertionSort) Sort(a []int) {
	size := len(a)
	if size < 2 {
		return
	}
	for i := 1; i < size; i++ {
		var j int
		var buff = a[i]
		for j = i - 1; j >= 0; j-- {
			if a[j] < buff {
				break
			}
			a[j+1] = a[j]
		}
		a[j+1] = buff
	}
}

// ContextOfSort - контекст
type ContextOfSort struct {
	strategy StrategySort
}

// Algorithm меняет стратегии
func (c *ContextOfSort) Algorithm(a StrategySort) {
	c.strategy = a
}

// Sort - сортирует данные способом в зависимости от стратегии
func (c *ContextOfSort) Sort(s []int) {
	c.strategy.Sort(s)
}

// клиентский код
func main() {
	data1 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}
	data2 := []int{8, 2, 6, 7, 1, 3, 9, 5, 4}

	ctx := new(ContextOfSort)

	ctx.Algorithm(&BubbleSort{})

	ctx.Sort(data1)

	ctx.Algorithm(&InsertionSort{})

	ctx.Sort(data2)

	fmt.Println(data1)
	fmt.Println(data2)
}