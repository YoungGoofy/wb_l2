package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Позволяет создавать объекты без указания конкретного класса.
Вместо этого клиентский код обращается к фабрике, которая возвращает нужный объект.
*/

const (
	audi = "audi"
	bmw  = "bmw"
)

type Info interface {
	PrintCar()
}

type Car1 struct {
	Color string
	Model string
}

func (c *Car1) PrintCar() {
	fmt.Printf("model: [%v] color: [%v]", c.Model, c.Color)
}

func NewCar1(model string, color string) *Car1 {
	return &Car1{Model: model, Color: color}
}

type Car2 struct {
	Color string
	Model string
}

func (c *Car2) PrintCar() {
	fmt.Printf("model: [%v] color: [%v]", c.Model, c.Color)
}

func NewCar2(model string, color string) *Car2 {
	return &Car2{Model: model, Color: color}
}

func New(model string) Info {
	switch model {
	case audi:
		return NewCar1("audi", "black")
	case bmw:
		return NewCar2("bmw", "pink")
	}
	return nil
}

func main() {
	car := New(audi)
	car.PrintCar()
}
