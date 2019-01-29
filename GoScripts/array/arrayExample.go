package main

import (
	"fmt"
)

func SubtractOneFromLength(slice []int) []int {
	slice = slice[0 : len(slice)-1]
	return slice
}

func main() {
	// arrays must be initialised with a fixed length
	var intArray [10]int
	intArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array length:", len(intArray))
	fmt.Println("5th element:", intArray[4])

	fmt.Println("arrays have a fixed length, slices are more dynamic")

	// slices aren't, but they do have a length based on how you fill them

	intSlice := []int{}
	fmt.Println("Slice length:", len(intSlice))

	intSlice = intArray[5:10]
	fmt.Println("Slice length:", len(intSlice))

	fmt.Println(intArray)
	fmt.Println(intSlice)

	// it is important to note that the slice is not a dynamic array, it is a dynamic representation of the array
	// hence updating the array affects the slice!
	fmt.Println("updating the array affects the slice!")
	intArray[9] = 11

	fmt.Println(intArray)
	fmt.Println(intSlice)

	fmt.Println("and vice versa!")
	intSlice[len(intSlice)-1] = 10
	fmt.Println(intArray)
	fmt.Println(intSlice)

	fmt.Println("if you need a copy, make a copy!")
	var copiedSl = make([]int, len(intSlice))
	copy(copiedSl, intSlice)
	copy(copiedSl, intSlice)
	copiedSl[0] = 1007

	fmt.Println(intArray)
	fmt.Println(intSlice)
	fmt.Println(copiedSl)

	fmt.Println("but what happens here?")
	fmt.Println(SubtractOneFromLength(intSlice))
	fmt.Println(intSlice)

	intSlice = SubtractOneFromLength(intSlice)
	fmt.Println(intSlice)
	fmt.Println(intArray)

	intSlice = append(intSlice, 73)

	fmt.Println("as long as the slice still can be represented by the array the connection stays alive")
	intArray[9] = 99
	fmt.Println(intArray)
	fmt.Println(intSlice)
	intArray[8] = 98
	fmt.Println(intArray)
	fmt.Println(intSlice)

	intSlice = append(intSlice, 73)
	intSlice = append(intSlice, 37)

	fmt.Println("now part of the slice points to a different array!")
	intArray[9] = 100
	fmt.Println(intArray)
	fmt.Println(intSlice)
	intArray[8] = 101
	fmt.Println(intArray)
	fmt.Println(intSlice)
}
