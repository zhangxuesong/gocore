package main

import "fmt"

func main() {
	s1 := make([]int, 5)
	fmt.Printf("the length of s1: %d\n", len(s1))
	fmt.Printf("the capacity of s1: %d\n", cap(s1))
	fmt.Printf("the value of s1: %d\n", s1)

	s2 := make([]int, 5, 8)
	fmt.Printf("the length of s2: %d\n", len(s2))
	fmt.Printf("the capacity of s2: %d\n", cap(s2))
	fmt.Printf("the value of s2: %d\n", s2)

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("the length of s4: %d\n", len(s4))
	fmt.Printf("the capacity of s4: %d\n", cap(s4))
	fmt.Printf("the value of s4: %d\n", s4)

	s5 := s4[0:cap(s4)]
	fmt.Printf("the length of s5: %d\n", len(s5))
	fmt.Printf("the capacity of s5: %d\n", cap(s5))
	fmt.Printf("the value of s5: %d\n", s5)

	s6 := make([]int, 0)
	fmt.Printf("The capacity of s6: %d\n", cap(s6))
	for i := 1; i <= 5; i++ {
		s6 = append(s6, i)
		fmt.Printf("s6(%d): len: %d, cap: %d\n", i, len(s6), cap(s6))
	}

	//The capacity of s6: 0
	//s6(1): len: 1, cap: 1
	//s6(2): len: 2, cap: 2
	//s6(3): len: 3, cap: 4
	//s6(4): len: 4, cap: 4
	//s6(5): len: 5, cap: 8

	fmt.Println()

	// 示例2。
	s7 := make([]int, 1024)
	fmt.Printf("The capacity of s7: %d\n", cap(s7))
	s7e1 := append(s7, make([]int, 200)...)
	fmt.Printf("s7e1: len: %d, cap: %d\n", len(s7e1), cap(s7e1))
	s7e2 := append(s7, make([]int, 400)...)
	fmt.Printf("s7e2: len: %d, cap: %d\n", len(s7e2), cap(s7e2))
	s7e3 := append(s7, make([]int, 600)...)
	fmt.Printf("s7e3: len: %d, cap: %d\n", len(s7e3), cap(s7e3))

	//The capacity of s7: 1024
	//s7e1: len: 1224, cap: 1280
	//s7e2: len: 1424, cap: 1696
	//s7e3: len: 1624, cap: 2048
	fmt.Println()

	// 示例3。
	s8 := make([]int, 10)
	fmt.Printf("The capacity of s8: %d\n", cap(s8))
	s8a := append(s8, make([]int, 11)...)
	fmt.Printf("s8a: len: %d, cap: %d\n", len(s8a), cap(s8a))
	s8b := append(s8a, make([]int, 23)...)
	fmt.Printf("s8b: len: %d, cap: %d\n", len(s8b), cap(s8b))
	s8c := append(s8b, make([]int, 45)...)
	fmt.Printf("s8c: len: %d, cap: %d\n", len(s8c), cap(s8c))

	//The capacity of s8: 10
	//s8a: len: 21, cap: 22
	//s8b: len: 44, cap: 44
	//s8c: len: 89, cap: 96
	fmt.Println()

	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1))
	s9 := a1[1:4]
	//s9[0] = 1
	fmt.Printf("s9: %v (len: %d, cap: %d)\n",
		s9, len(s9), cap(s9))
	for i := 1; i <= 5; i++ {
		s9 = append(s9, i)
		fmt.Printf("s9(%d): %v (len: %d, cap: %d)\n",
			i, s9, len(s9), cap(s9))
	}
	fmt.Printf("a1: %v (len: %d, cap: %d)\n",
		a1, len(a1), cap(a1))

	//a1: [1 2 3 4 5 6 7] (len: 7, cap: 7)
	//s9: [2 3 4] (len: 3, cap: 6)
	//s9(1): [2 3 4 1] (len: 4, cap: 6)
	//s9(2): [2 3 4 1 2] (len: 5, cap: 6)
	//s9(3): [2 3 4 1 2 3] (len: 6, cap: 6)
	//s9(4): [2 3 4 1 2 3 4] (len: 7, cap: 12)
	//s9(5): [2 3 4 1 2 3 4 5] (len: 8, cap: 12)
	//a1: [1 2 3 4 1 2 3] (len: 7, cap: 7)

	fmt.Println()
}
