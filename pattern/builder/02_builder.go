package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Паттерн Builder - порождающий паттерн проектирования, который позволяет пошагово создавать сложные объекты
Он позволяет создать объекты с различными св-вами и конфигами не засоряя конструктор
*/

type Car struct {
	Model string
	Color string
}

type Builder interface {
	SetModel(string) Builder
	SetColor(string) Builder
	Build() Car
}

type carBuilder struct {
	model string
	color string
}

func (c *carBuilder) SetModel(model string) Builder {
	c.model = model
	return c
}

func (c *carBuilder) SetColor(color string) Builder {
	c.color = color
	return c
}

func (c *carBuilder) Build() Car {
	return Car{
		Model: c.model,
		Color: c.color,
	}
}

func main() {
	builder := carBuilder{}
	car := builder.SetColor("green").SetModel("bmw").Build()
	fmt.Println(car)
}
