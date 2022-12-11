package main

import "fmt"

func main() {
	fmt.Println("main function is called")
	fmt.Println("-------------------------------------")
	// To declare and use the constants
	Const()

	fmt.Println("-------------------------------------")
	// To identify the scenarios of data overflow
	Overflow()
	
	// Demonstrate the Size of the memory in KB,MB,GB,TB using enum
	MemorySize()
}
