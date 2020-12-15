package main

import (
	"errors"
	"fmt"
)

type operates func(x, y int) int

type calculator func(x, y int) (int, error)

func genCalculator(op operates) calculator {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("无效操作")
		}
		return op(x, y), nil
	}
}

func main() {
	op := func(x, y int) int {
		return x + y
	}

	x, y := 56, 78
	add := genCalculator(op)
	result, err := add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}
