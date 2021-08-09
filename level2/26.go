package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	// Создается новый локальный слайс, тк в старом кончилось cap
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
	// локальный i умирает
}

// [3 2 3]
