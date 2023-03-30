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

/*
Используется для обхода объектов в структуре данных и выполнения определенных операций над каждым из объектов.
При этом паттерн позволяет добавлять новые операции, не изменяя классы объектов, над которыми они выполняются.
*/

type Shape interface {
	Accept(Visitor)
}

type Circle struct {
	radius float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

func NewCircle(radius float64) Shape {
	return &Circle{radius: radius}
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitSquare(r)
}

func NewRectangle(height float64, width float64) Shape {
	return &Rectangle{height: height, width: width}
}

type Visitor interface {
	VisitCircle(*Circle)
	VisitSquare(*Rectangle)
}

type AreaVisitor struct {
	area float64
}

func (av *AreaVisitor) VisitCircle(c *Circle) {
	av.area += math.Pi * c.radius * c.radius
}

func (av *AreaVisitor) VisitSquare(r *Rectangle) {
	av.area += r.height * r.width
}

func main() {
	area := AreaVisitor{}

	circle := NewCircle(5)
	circle.Accept(&area)

	rectangle := NewRectangle(2, 3)
	rectangle.Accept(&area)

	fmt.Println(area.area)
}
