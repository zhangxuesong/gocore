package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

func calculate(x, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("无效操作")
	}
	return op(x, y), nil
}

func main() {
	op := func(x, y int) int {
		return x + y
	}
	s, err := calculate(1, 2, op)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(s)
}
