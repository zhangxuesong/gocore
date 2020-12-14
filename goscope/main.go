package main

import "fmt"

var block = "package"

func main() {
	block := "function"
	{
		block := "inner"
		fmt.Printf("this block is %s!\n", block)
	}
	fmt.Printf("this block is %s!\n", block)
}
