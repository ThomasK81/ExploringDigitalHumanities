package main

import (
	"fmt"
)

func main() {
	newString := "test"
	fmt.Println(newString)
	fmt.Println(newString[0:2])

	fmt.Println(newString[0])

	fmt.Println([]byte(newString))

	fmt.Println(string(newString[0]))

	newString = "τεστ"
	fmt.Println(newString)
	fmt.Println(newString[0:2])
	fmt.Println(newString[0])

	fmt.Println(string(newString[0]))
	fmt.Println([]byte(newString))

	fmt.Println([]rune(newString))

	fmt.Println(string([]rune(newString)[0]))

}
