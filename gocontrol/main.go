package main

import "fmt"

func main() {
	//numbers := []int{1, 2, 3, 4, 5, 6}
	//for i := range numbers{
	//	if i == 3 {
	//		numbers[i] |= i
	//	}
	//}
	//fmt.Println(numbers)

	//numbers := [...]int{1, 2, 3, 4, 5, 6}
	//maxIndex := len(numbers) - 1
	//for k, v := range numbers {
	//	if k == maxIndex {
	//		numbers[0] += v
	//	} else {
	//		numbers[k+1] += v
	//	}
	//}
	//fmt.Println(numbers)

	numbers := []int{1, 2, 3, 4, 5, 6}
	maxIndex := len(numbers) - 1
	for k, v := range numbers {
		if k == maxIndex {
			numbers[0] += v
		} else {
			numbers[k+1] += v
		}
	}
	fmt.Println(numbers)
}
