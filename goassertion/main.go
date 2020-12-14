package main

import "fmt"

var container = []string{"0", "1", "2"}

func main() {
	container := map[int]string{0: "0", 1: "1", 2: "2"}

	_, ok1 := interface{}(container).([]string)
	_, ok2 := interface{}(container).(map[int]string)
	if !ok1 && !ok2 {
		fmt.Printf("Error: unsupported container type: %T\n", container)
		return
	}
	fmt.Printf("the element is %q.(container type: %T)\n", container[1], container)
	var srcInt = int16(-255)
	dstInt := int8(srcInt)
	fmt.Println(dstInt)

	fmt.Println(string(-1))

	fmt.Println(string([]byte{'\xe4', '\xbd', '\xa0', '\xe5', '\xa5', '\xbd'}))
}
