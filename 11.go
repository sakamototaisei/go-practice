package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("hello " + "world")
	fmt.Println(string("hello world"[0]))

	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "World"))

	fmt.Println("Test")
}
