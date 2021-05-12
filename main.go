package main

import "fmt"

var score = 99.5

// cannot use shorthand outside of functions
// scoreTwo := 50

func main() {
	sayHello("mario")
	showScore()

	for _, v := range points {
		fmt.Println(v)
	}
}
