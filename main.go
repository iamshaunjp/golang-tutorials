package main

import (
	"fmt"
)

func main() {
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++ // infinite loop without this
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, val := range names {
		fmt.Printf("the value at pos %v is %v \n", index, val)
		val = "new string"
	}

	for _, val := range names {
		fmt.Print(val, ",")
		val = "new string"
	}

	// changing val in a loop does not mutate the original slice
	fmt.Println(names)

}
