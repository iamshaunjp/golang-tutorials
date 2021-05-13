package main

import "fmt"

// func updateName(x string) {
// 	x = "yoshi"
// }

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// value types -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	// updateName(name)
	name = updateName(name)

	fmt.Println(name)

	// "reference" types -> slices, maps, functions, pointers
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
