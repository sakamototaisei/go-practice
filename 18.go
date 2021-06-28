package main

import (
	"fmt"
)

func main() {
	b := []byte{83, 65, 75, 65, 84, 65, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("SAKATAI")
	fmt.Println(c)
	fmt.Println(string(c))
}
