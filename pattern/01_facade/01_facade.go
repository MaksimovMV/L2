package main

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

//фасад предоставляет простой (но урезанный) итерфейс к сложной системе
//применимость: необходимость предоставить простой или урезанный интерфейс к сложной подсистеме
//преимущества: изолирует клиентов от компонентов сложной подсистемы
//недостатки: фасад рискует стать "божественным объектом", привязанным ко всем классам подпрограммы

// Computer  - фасад, предоставляет быстрый доступ к определенной функциональности системы
type Computer struct {
	cpu       *cpu
	memory    *memory
	hardDrive *hardDrive
}

func newComputer() *Computer {
	return &Computer{
		cpu:       newCPU(),
		memory:    newMemory(),
		hardDrive: newHardDrive(),
	}
}

func (c *Computer) startComputer() {
	c.cpu.freeze()
	lba := "100"
	size := "1024"
	ssd := c.hardDrive.read(lba, size)
	position := "0x00"
	c.memory.load(position, ssd)
	c.cpu.jump(position)
	c.cpu.execute()
}

//часть сложной подсистемы
type cpu struct {
}

func newCPU() *cpu {
	return &cpu{}
}

func (c *cpu) freeze() {
	fmt.Println("Freezing processor")
}

func (c *cpu) jump(position string) {
	fmt.Printf("Jumping to %v\n", position)
}

func (c *cpu) execute() {
	fmt.Println("Executing")
}

//часть сложной подсистемы
type memory struct {
}

func newMemory() *memory {
	return &memory{}
}

func (m *memory) load(position, data string) {
	fmt.Printf("Loading from %v data: %v\n", position, data)
}

//часть сложной подсистемы
type hardDrive struct {
}

func newHardDrive() *hardDrive {
	return &hardDrive{}
}

func (h *hardDrive) read(lba, size string) string {
	fmt.Println("reading")
	return fmt.Sprintf("Some data from sector %v with size %v", lba, size)
}

//клиентский код
func main() {
	pc := newComputer()
	pc.startComputer()
}
