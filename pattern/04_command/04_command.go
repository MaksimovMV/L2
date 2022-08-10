package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

//команда - поведенческий паттерн, позволяющий использовать функции как объекты (добавлять в очереди, удалять и т.д.)
//применение: 1) необходимость параметризировать объекты выполняемым действием
// 2) необходимость ставить операции в очередь
// 3) необходимость отмены операции
// плюсы: 1) нет прямой зависимости между объектами, вызывающими операцию, и
//объектами, которые их непосредственно выполняют
//2) возможность реализовать отмену и повтор операций
//3) возможность реализовать отложенный запуск операций
//4) возможность собирать сложные команды из простых
//5) реализует принцип открытости/закрытости
// минусы: усложняет код программы из-за введения множества дополнительных классов

// Command - интерфейс команды (обычно один метод для запуска)
type Command interface {
	Execute()
}

// OnCommand - Конкретная команда
type OnCommand struct {
	receiver Receiver
}

func (c *OnCommand) Execute() {
	c.receiver.on()
}

// OffCommand - Конкретная команда
type OffCommand struct {
	receiver Receiver
}

func (c *OffCommand) Execute() {
	c.receiver.off()
}

// Receiver - Интерфейс получателя, содержит бизнес-логику программы, команды перенаправляют вызов получателям
type Receiver interface {
	on()
	off()
}

// Computer - Конкретный получатель
type Computer struct {
}

func (c *Computer) on() {
	fmt.Println("Computer On")
}

func (c *Computer) off() {
	fmt.Println("Computer Off")
}

// Invoker - Инициатор/Отправитель - хранит команды, работает с ними только через их интерфейс
type Invoker struct {
	commands []Command
}

// StoreCommand - Добавление команд
func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

// UnStoreCommand - Удаление команд
func (i *Invoker) UnStoreCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// ExecuteAll - Запуск всех команд
func (i *Invoker) ExecuteAll() {
	for _, command := range i.commands {
		command.Execute()
	}

}

// Пользовательский код
func main() {
	comp := &Computer{}

	onCommand := &OnCommand{
		receiver: comp,
	}

	offCommand := &OffCommand{
		receiver: comp,
	}

	inv := &Invoker{
		commands: make([]Command, 0),
	}

	for i := 0; i < 4; i++ {
		inv.StoreCommand(onCommand)
		inv.StoreCommand(offCommand)
	}

	inv.UnStoreCommand()

	inv.ExecuteAll()
}
