package main

import "fmt"

//only for loops available in Golang
func main() {
	//loop from 0 until 5 is reached
	for i := 1; i <= 5; i++ {
		fmt.Println("go=", i)
	}

	list := []int{1, 3, 45, 13, 59, 0}
	//loop itarates for a range over a slice or map.
	for z, v := range list {
		fmt.Printf("Element %d = %d\n", z, v)
	}
}

