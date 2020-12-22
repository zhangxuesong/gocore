package main

import (
	"errors"
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "Joseph", "the greeting object.")
}

func main() {
	flag.Parse()
	greeting, err := hello(name)
	if err != nil {
		fmt.Printf("error is: %s\n", err)
		return
	}
	fmt.Println(greeting, introduce())
}

func hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name.")
	}
	return fmt.Sprintf("hello %s!", name), nil
}

func introduce() string {
	return "welcome to my golang column"
}
