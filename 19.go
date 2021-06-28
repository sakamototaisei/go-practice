package main

import "fmt"

func add(x int, y int) (int, int) {
	// fmt.Println("add function")
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(200, 3)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int) {
		fmt.Println("inner func", x)
	}(1)
}
