package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Позволяет определять набор алгоритмов, инкапсулировать каждый из них и делать их взаимозаменяемыми.
При этом позволяется изменять алгоритмы независимо от клиентского кода, который их использует.
*/

type PaymentMethod interface {
	Pay(amount float64) string
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid from credit card", amount)
}

type Cash struct{}

func (c *Cash) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f paid from cash", amount)
}

type Product struct {
	name   string
	amount float64
	method PaymentMethod
}

func (p *Product) Pay() string {
	return p.method.Pay(p.amount)
}

func (p *Product) SetPaymentMethod(method PaymentMethod) {
	p.method = method
}

func main() {
	product1 := &Product{name: "Banana", amount: 12.5, method: &Cash{}}
	product2 := &Product{name: "Car", amount: 15.5, method: &CreditCard{}}

	fmt.Println(product1.Pay())
	fmt.Println(product2.Pay())

	product1.SetPaymentMethod(&CreditCard{})
	fmt.Println(product1.Pay())
}
