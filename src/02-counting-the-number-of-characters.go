package main

import "fmt"

func main() {
	var str string

	fmt.Println("What is the input string?")
	fmt.Scan(&str)
	fmt.Printf(str+" has %v characters\n", len(str))
}
