package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}

/*
делаем слайс от 1 до 4 из массва элемента и выводим
[77 78 79]

*/
