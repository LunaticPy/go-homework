package main

import (
	"fmt"
)

func test() (x int) {
	// x= 0
	defer func() {
		x++
	}()

	// x=1
	x = 1
	return
	// defer x=2
}

func anotherTest() int {
	var x int
	// x инициализирован, но не имеет значения
	defer func() {
		// у х нет значения
		x++
	}()

	// x=1
	x = 1
	return x
	// defer дефер "ничего не делает"

}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

// 2
// 1
