package main

import (
	"flag"
	"fmt"
	"os"
)

//var name = flag.String("name", "Joseph", "The greeting object.")
var name string
var cmdLine = flag.NewFlagSet("参数", flag.ExitOnError)

func init() {
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	//flag.CommandLine.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "%s使用说明:\n", "参数")
	//	flag.PrintDefaults()
	//}

	//flag.StringVar(&name, "name", "Joseph", "The greeting object.")

	cmdLine.StringVar(&name, "name", "Joseph", "The greeting object.")
}

func main()  {
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "%s使用说明:\n", "参数")
	//	flag.PrintDefaults()
	//}
	//flag.Parse()

	cmdLine.Parse(os.Args[1:])
	fmt.Printf("hello %s!\n", name)
}
