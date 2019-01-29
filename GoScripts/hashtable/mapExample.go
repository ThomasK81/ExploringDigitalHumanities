package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "Hello"
	m[2] = "World"
	fmt.Println(m[1], m[2])

	var n map[string][]int
	n = make(map[string][]int)

	n["one"] = []int{1, 2, 4}
	n["two"] = []int{5, 6, 8}
	fmt.Println(n["one"], n["two"])

	o := make(map[string]int)
	fmt.Println(n["not"])
	fmt.Println(m[3])
	fmt.Println(o["not"])

	delete(n, "two")

	i, ok := n["two"]
	if ok {
		fmt.Println(i)
	} else {
		fmt.Println("not found")
	}

	i, ok = n["one"]
	if ok {
		fmt.Println(i)
	} else {
		fmt.Println("not found")
	}

	// remember for later: maps are not safe for concurrency
	// remember for later: iteration does not work always in the same order
}
