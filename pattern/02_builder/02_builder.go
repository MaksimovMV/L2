package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

//строитель позволяет создавать объекты пошагово
//применение: 1) когда код должен создавать разные представления какого-то объекта
// 2) когда нужно собирать сложные составные объекты
// плюсы: 1) позволяет создавать продукты пошагово
// 2) позволяет использовать один и тот же код для создания различных продуктов
// 3) изолирует сложный код сборки продукта от его основной бизнес-логики
// минусы: 1) усложняет код программы из-за введения дополнительных классов
// 2) клиент будет привязан к конкретным класссам строителей, т.к. в интерфейсе
//директора может не быть метода получения результата

// продукт (не обязан после создания иметь общий интерфейс)
type pizza struct {
	dough   string
	sauce   string
	topping string
}

// интерфейс строителя, объявляет общие шаги конструирования продуктов
type iBuilder interface {
	setDough()
	setSauce()
	setTopping()
	getPizza() pizza
}

func getBuilder(builderType string) iBuilder {
	if builderType == "hawaiian" {
		return newHawaiianPizzaBuilder()
	}

	if builderType == "spicy" {
		return newSpicyPizzaBuilder()
	}

	return nil
}

// конкретный строитель
type hawaiianPizzaBuilder struct {
	dough   string
	sauce   string
	topping string
}

func newHawaiianPizzaBuilder() *hawaiianPizzaBuilder {
	return &hawaiianPizzaBuilder{}
}

func (b *hawaiianPizzaBuilder) setDough() {
	b.dough = "cross"
}

func (b *hawaiianPizzaBuilder) setSauce() {
	b.sauce = "mild"
}

func (b *hawaiianPizzaBuilder) setTopping() {
	b.topping = "ham+pineapple"
}
func (b *hawaiianPizzaBuilder) getPizza() pizza {
	return pizza{
		dough:   b.dough,
		sauce:   b.sauce,
		topping: b.topping,
	}
}

// конкретный строитель
type spicyPizzaBuilder struct {
	dough   string
	sauce   string
	topping string
}

func newSpicyPizzaBuilder() *spicyPizzaBuilder {
	return &spicyPizzaBuilder{}
}

func (b *spicyPizzaBuilder) setDough() {
	b.dough = "pan baked"
}

func (b *spicyPizzaBuilder) setSauce() {
	b.sauce = "hot"
}

func (b *spicyPizzaBuilder) setTopping() {
	b.topping = "pepperoni+salami"
}
func (b *spicyPizzaBuilder) getPizza() pizza {
	return pizza{
		dough:   b.dough,
		sauce:   b.sauce,
		topping: b.topping,
	}
}

// Директор, определяет порядок вызова шагов
type waiter struct {
	builder iBuilder
}

func newWaiter(b iBuilder) *waiter {
	return &waiter{
		builder: b,
	}
}

func (w *waiter) setBuilder(b iBuilder) {
	w.builder = b
}

func (w *waiter) buildPizza() pizza {
	w.builder.setDough()
	w.builder.setSauce()
	w.builder.setTopping()
	return w.builder.getPizza()
}

// клиентский код
func main() {
	hawaiianBuilder := getBuilder("hawaiian")
	spicyBuilder := getBuilder("spicy")

	waiter := newWaiter(hawaiianBuilder)
	hawaiianPizza := waiter.buildPizza()

	fmt.Printf("Hawaiian Pizza dough: %s\n", hawaiianPizza.dough)
	fmt.Printf("Hawaiian Pizza sauce: %s\n", hawaiianPizza.sauce)
	fmt.Printf("Hawaiian Pizza topping: %s\n", hawaiianPizza.topping)

	waiter.setBuilder(spicyBuilder)
	spicyPizza := waiter.buildPizza()

	fmt.Printf("Spicy Pizza dough: %s\n", spicyPizza.dough)
	fmt.Printf("Spicy Pizza sauce: %s\n", spicyPizza.sauce)
	fmt.Printf("Spicy Pizza topping: %s\n", spicyPizza.topping)
}