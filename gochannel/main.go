package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//ch := make(chan int, 3)
	//ch <- 2
	//ch <- 1
	//ch <- 3
	//elmt, ok := <- ch
	//fmt.Println(ok)
	//fmt.Printf("the first element is: %v.\n", elmt)

	//ch1 := make(chan<- int)//发送通道，只能发不能收
	//ch2 := make(<-chan int)//接收通道，只能收不能发

	//ch3 := make(chan int, 1)
	//SendInt(ch3)
	//elmt3, ok := <- ch3
	//fmt.Printf("elemt3 is %v. channel stats is %v.\n", elmt3, ok)

	//ch4 := getIntChan()
	//for elemt := range ch4 {
	//	fmt.Printf("the element in ch4: %v.\n", elemt)
	//}

	//// 准备好几个通道。
	//intChannels := [3]chan int{
	//	make(chan int, 1),
	//	make(chan int, 1),
	//	make(chan int, 1),
	//}
	//// 随机选择一个通道，并向它发送元素值。
	//index := rand.Intn(3)
	//fmt.Printf("The index: %d\n", index)
	//intChannels[index] <- index
	//// 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
	//select {
	//case <-intChannels[0]:
	//	fmt.Println("The first candidate case is selected.")
	//case <-intChannels[1]:
	//	fmt.Println("The second candidate case is selected.")
	//case elem := <-intChannels[2]:
	//	fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	//default:
	//	fmt.Println("No candidate case is selected!")
	//}

	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
}

func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
	//fmt.Printf("element of channel: %v.\n", <- ch)
}

func getIntChan() <-chan int {
	num := 5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}
