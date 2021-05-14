package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")

	fmt.Println(mybill.format())
}
