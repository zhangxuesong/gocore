package main

import (
	"flag"
	lib5 "gocore/libSource/lib"
	//"libsource/lib/internal"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Joseph", "The greeting object.")
}

func main() {
	flag.Parse()
	//hello(name)
	lib5.Hello(name)
}
