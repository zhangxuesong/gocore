package main

import "fmt"

func main() {
	//aMap := map[string]int{
	//	"one":   1,
	//	"two":   2,
	//	"three": 3,
	//}
	//k := "two"
	//v, ok := aMap[k]
	//if ok {
	//	fmt.Printf("the element of key %q: %d.\n", k, v)
	//} else {
	//	fmt.Println("not found!")
	//}

	//bMap := map[interface{}]int{
	//	"one":    1,
	//	[]int{2}: 2,
	//	3:        3,
	//}
	//fmt.Println(bMap)

	//cMap := map[[1][]string][]string
	//_ = cMap

	var dMap map[string]int

	v, ok := dMap["one"]
	fmt.Printf("The element paired in nil map: %d (%v)\n", v, ok)

	dMap["one"] = 1
	fmt.Println(dMap)
}
