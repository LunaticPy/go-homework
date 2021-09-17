package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil

	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	// type=nil, val=nil != type=*os.PathError val=nil
	fmt.Println(err == nil)
}

// <nil>
// false
