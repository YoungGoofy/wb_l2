package main

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Позволяет изменять поведение объекта в зависимости от его внутреннего состояния.
Этот шаблон основан на принципе разделения ответственности и помогает облегчить расширение и поддержку кода
*/

import "fmt"

type Context struct {
	state State
}

func (c *Context) setState(state State) {
	fmt.Println("Context: изменение состояния.")
	c.state = state
}

func (c *Context) request() {
	c.state.handle()
}

type State interface {
	handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) handle() {
	fmt.Println("ConcreteStateA: обработка запроса.")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) handle() {
	fmt.Println("ConcreteStateB: обработка запроса.")
}

func main() {
	context := &Context{}

	stateA := &ConcreteStateA{}
	context.setState(stateA)
	context.request()

	stateB := &ConcreteStateB{}
	context.setState(stateB)
	context.request()
}
