package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Способен организовать обработку запроса последовательным прохождением по цепочке объектов, которые могут
обрабатывать запрос или передавать его дальше по цепочке
*/

type Request struct {
	content string
}

type Handler interface {
	SetNext(handler Handler)
	Handle(request Request)
}

type AbstractHandler struct {
	next Handler
}

func (a *AbstractHandler) SetNext(handler Handler) {
	a.next = handler
}

type ConcreteHandler1 struct {
	AbstractHandler
}

func (h *ConcreteHandler1) Handle(request Request) {
	if request.content == "request1" {
		fmt.Println("request1 start")
	} else if h.next != nil {
		h.next.Handle(request)
	}
}

type ConcreteHandler2 struct {
	AbstractHandler
}

func (h *ConcreteHandler2) Handle(request Request) {
	if request.content == "request2" {
		fmt.Println("request2 start")
	} else if h.next != nil {
		h.next.Handle(request)
	}
}

type ConcreteHandler3 struct {
	AbstractHandler
}

func (h *ConcreteHandler3) Handle(request Request) {
	if request.content == "request1" {
		fmt.Println("request3 start")
	} else if h.next != nil {
		h.next.Handle(request)
	}
}

func main() {
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}
	handler3 := &ConcreteHandler3{}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	requests := []Request{
		Request{"request1"},
		Request{"request2"},
		Request{"request3"},
		Request{"request4"},
	}

	for _, request := range requests {
		handler1.Handle(request)
	}
}
