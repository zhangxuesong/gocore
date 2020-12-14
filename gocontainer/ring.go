package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.Ring{}

	fmt.Printf("the ring is: %v.\n", r.Len())
}
