package main

import "fmt"

// func updateName(n string) {
// 	n = "wedge"
// }

func updateName(n *string) {
	*n = "wedge"
}

func main() {
	// value types -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	// updateName(name)
	// fmt.Println(name)

	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateName(&name) // using pointer as arg
	fmt.Println(name)

}
