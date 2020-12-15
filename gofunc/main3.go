package main

import "fmt"

func main() {
	array1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
}

func modifyArray(a [3][]string) [3][]string {
	a[0][1] = "z"
	a[1] = []string{"a", "b", "c"}
	return a
}
