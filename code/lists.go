package main

import "fmt"

func main() {
	var arr [3]int                             //initiate list [0.0.0]
	arr[0] = 4                                 //assign value to element
	colors := []string{"blue", "black", "red"} //SLICES - when you don't know a size of a list
	colors = append(colors, "orange")          //append == recreation

	// length, capacity
	fmt.Println(len(colors))
	fmt.Println(cap(colors))

	// show results
	fmt.Println("slice 'colors': ", colors)
	fmt.Println("list 'arr': ", arr)
}
