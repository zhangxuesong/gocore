package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println(s[5])
}

//panic: runtime error: index out of range [5] with length 5
//
//goroutine 1 [running]:
//main.main()
///Users/zhangxuesong/gowork/src/gocore/gopanic/main.go:7 +0x1b
//exit status 2
