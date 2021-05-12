package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("enter your name: ")
	// name, _ := reader.ReadString('\n')
	// fmt.Println("your name is", input)

	opt, _ := getInput("Choose an option (a -add item, s - save bill, t - add tip): ", reader)
	fmt.Println(opt)
}

func main() {

	mybill := newBill("mario's bill")
	promptOptions(mybill)

}
