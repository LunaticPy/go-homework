package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	// по очереди читаем из a,b
	//  после окончания данных бесконечно читаем 0(нулевое значение)
	// тк при чтении range из канала бесконечно будет идти 0, false тк канал не закрыт и больше не получает данные, но false мы никак не обрабатываем
	// Если мы не закроем канал для цикла for с использованием range, то программа будет завершена аварийно из-за dealock во время выполнения.
	for v := range c {
		fmt.Println(v)
	}
}

// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// {0}
