package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

//Фабричный метод — это порождающий паттерн проектирования,
//который решает проблему создания различных продуктов, без указания
//конкретных классов продуктов.
// Применение 1) когда неизвестны типы и зависимости объектов, с которыми должен работать код
// 2) необходимо дать возможность расширить части вашего фреймворка или библиотеки
// 3) когда хотите экономить системные ресурсы, используя уже созданные объекты
// плюсы 1) избавляет класс от привязки к конкретным классам продуктов
// 2) выделят код производства продуктов в одно место, упрощая поддержку кода
// 3) упрощает добавление новых продуктов в программу
// 4) реализует принцип открытости/закрытости
// минус - модет привести к созданию больших параллельных иерархий классов,
// т.к. для каждого класса продукта надо создать свой класс создателя
// в Go нет возможности реализовать коассический вариант фабричного метода. Следующий
//пример - простая фабрика

// Product - интерфейс продукта
type Product interface {
	setName(name string)
	getName() string
}

// ProductBase - конкретный продукт
type ProductBase struct {
	name string
}

func (p *ProductBase) setName(name string) {
	p.name = name
}

func (p *ProductBase) getName() string {
	return p.name
}

// ProductA - конкретный продукт
type ProductA struct {
	ProductBase
}

func newProductA() Product {
	return &ProductA{
		ProductBase: ProductBase{
			name: "Product A",
		},
	}
}

// ProductB - конкретный продукт
type ProductB struct {
	ProductBase
}

func newProductB() Product {
	return &ProductB{
		ProductBase: ProductBase{
			name: "Product B",
		},
	}
}

// getProduct - Фабрика
func getProduct(productType string) (Product, error) {
	if productType == "A" {
		return newProductA(), nil
	}
	if productType == "B" {
		return newProductB(), nil
	}
	return nil, fmt.Errorf("wrong product type passed")
}

// клиенский код
func main() {
	a, _ := getProduct("A")
	b, _ := getProduct("B")
	fmt.Printf("Product: %s\n", a.getName())
	fmt.Printf("Product: %s\n", b.getName())
}
