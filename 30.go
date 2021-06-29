package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./hello.go")
	if err != nil {
		log.Fatalln("Error1")
	}
	defer file.Close()
	data := make([]byte, 200)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error2")
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil {
		log.Fatalln("Error3")
	}
}
