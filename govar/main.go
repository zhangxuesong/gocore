package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = getFlag()
	flag.Parse()
	fmt.Printf("Hello %v!\n", *name)
}

func getFlag() *string {
	return flag.String("name", "Joseph", "This greeting object.")
}
