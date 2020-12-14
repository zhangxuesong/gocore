package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.List{}
	fmt.Printf("the list init is: %v.\n", l)
	l.PushBack("hello")
	l.PushBack("world")
	fmt.Printf("the list is: %v.\n", l)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("the element is: %v.\n", e)
	}

	//the list is: {{0xc0000621b0 0xc0000621e0 <nil> <nil>} 2}.
	//the element is: &{0xc0000621e0 0xc000062180 0xc000062180 hello}.
	//the element is: &{0xc000062180 0xc0000621b0 0xc000062180 world}.

}
