package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")

	p := recover()
	// 引发panic。
	panic(errors.New("something wrong"))
	fmt.Printf("panic: %s\n", p)
	fmt.Println("Exit function main.")
}
