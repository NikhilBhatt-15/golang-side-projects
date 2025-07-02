package main

import "fmt"

func main() {
	var ptr *int
	var value = 42
	ptr = &value 
	*ptr = 100 
	fmt.Println("Value:", value) // Output: Value: 100
}