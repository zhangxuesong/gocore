package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	//
	//time.Sleep(time.Millisecond * 500)

	//num := 10
	//sign := make(chan struct{}, num)
	//
	//for i := 0; i < num; i++ {
	//	go func() {
	//		fmt.Println(i)
	//		sign <- struct{}{}
	//	}()
	//}
	//
	//for j := 0; j < num; j++ {
	//	<-sign
	//}

	var count uint32
	tigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			} else {
				time.Sleep(time.Nanosecond)
			}
		}
	}

	for i := uint32(0); i < 10; i++ {
		fn := func() {
			fmt.Println(i)
		}
		tigger(i, fn)
	}

	tigger(10, func() {})
}
