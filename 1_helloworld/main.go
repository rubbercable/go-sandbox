package main

import "fmt"

func main() {
	firstname := "Arthur"
	fmt.Println(firstname)

	ptr := &firstname
	fmt.Println(ptr, *ptr)

	firstname = "Tricia"
	fmt.Println(ptr, *ptr)
}
