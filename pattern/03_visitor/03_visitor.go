package main

import (
	"fmt"
	"math"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// посетитель - поведенческий паттерн, который позволяет добавить новую операцию для
// целой иерархии классов, не изменяя код этих классов
// применение: 1) применить одну и ту же операцию к объектам различных классов
// 2) новое поведение только для некоторых классов
// 3) извлечь родственные операции из классов и поместить в один класс посетитель
// плюсы: 1) упрощает добавление операций, работающих со сложными структурами объектов
// 2) объединяет родственные операции в одном классе
// 3) посетитель может накапливать состояние при обходе структур элементов
// минусы: 1) паттерн не оправдан, если иерархия постоянно меняется
// 2) может привести к нарушению инкапсуляции объектов

// Shape элемент
type Shape interface {
	getType() string
	accept(Visitor) //единственное изменение в интерфейсе и стуктурах
}

// Square - конкретный элемент
type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

// Circle - конкретный элемент
type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

// Rectangle - конкретный элемент
type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

// Visitor - посетитель
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// AreaCalculator - конкретный посетитель
type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	a.area = float64(s.side * s.side)

}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	a.area = float64(s.radius) * float64(s.radius) * math.Pi
}
func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	a.area = float64(s.l * s.b)
}

//ползовательский код
func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	fmt.Printf("Square area = %v\n", areaCalculator.area)

	circle.accept(areaCalculator)
	fmt.Printf("Circle area = %v\n", areaCalculator.area)

	rectangle.accept(areaCalculator)
	fmt.Printf("Rectangle area = %v\n", areaCalculator.area)
}
