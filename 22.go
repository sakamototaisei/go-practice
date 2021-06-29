package main

import "fmt"

func main() {
	f := 1.11
	i := int(f)
	fmt.Printf("%T %v\n", i, i)
	sub2()
	sub3()
}

func sub2() {
	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])
}

func sub3() {
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Println(m)
	fmt.Println(m["Mike"])
	fmt.Printf("%T %v\n", m, m)

	m2 := make(map[string]int)
	m2["Sakatai"] = 25
	fmt.Println(m2)
}
