package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("fkjgfdi")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test")
	log.Fatalln("error!")
}
