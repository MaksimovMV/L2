package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Цепочка обязанностей - поведенческий паттерн, который позволяет передавать запросы
//последовательно по цепочке обработчиков
// Применение - 1) необходимость обработки разнообразных запросов несколькими способами
// 2) необходимость выполнение обработчиков в строгом порядке
// 3) нужна возможность вмешаться в цепочку и переназначить связи
// Плюсы 1) уменьшает зависимость между клиентом и обработчиком
// 2) принцип единственной ответственности
// 3) принцип открытости/закрытости
// Минус - запрос может остаться никем не обработанным

// Handler - Интерфейс обработчика (обычно с одним методом обработки запроса)
type Handler interface {
	SendRequest(message int) string
}

// ConcreteHandlerA - Конкретный обработчик А
type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "Im handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerB - Конкретный обработчик В
type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "Im handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

// ConcreteHandlerC - Конкретный обработчик С
type ConcreteHandlerC struct {
	next Handler
}

func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "Im handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

//клиентский код
func main() {
	handlers := &ConcreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}

	fmt.Println(handlers.SendRequest(1))
	fmt.Println(handlers.SendRequest(3))
	fmt.Println(handlers.SendRequest(2))

}
