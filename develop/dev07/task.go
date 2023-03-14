package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {
	sig := func(after time.Duration, counter int) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			fmt.Printf("Горутина №%d начала работу\n", counter)
			time.Sleep(after)
			fmt.Printf("Горутина №%d закончила работу\n", counter)

		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(5*time.Second, 1),
		sig(500*time.Millisecond, 2),
		sig(1*time.Second, 3),
		sig(10*time.Second, 4),
		sig(250*time.Millisecond, 5),
	)

	fmt.Printf("Done after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})

	// Функция, которая отправляет значение из канала в out и удаляет его из списка channels
	send := func(ch <-chan interface{}) {

		for val := range ch {
			select {
			case out <- val:
			default:
			}
		}
		wg.Done()
	}

	// Добавляем каждый канал в wait group и запускаем go-рутину для отправки значений
	wg.Add(len(channels))
	for _, ch := range channels {
		go send(ch)
	}

	// Функция, которая закрывает канал out после того, как все каналы в wait group закроются
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
