package main

import (
	"fmt"
	"os/user"
	"time"
)

/*
func init() {
	fmt.Println("init")
}
*/
func bazz() {
	fmt.Println("Bazz")
}
func main() {
	//bazz()
	fmt.Println("hello,world", "test test", time.Now())
	fmt.Printf(user.Current())
}
