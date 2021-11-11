package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Print("What is the first number? ")
	first := read()
	first_i := to_int(first)

	fmt.Print("What is the second number? ")
	second := read()
	second_i := to_int(second)

	result_add := first_i + second_i
	result_sub := first_i - second_i
	result_mul := first_i * second_i
	result_div := first_i / second_i

	result_add_ := strconv.Itoa(result_add)
	result_sub_ := strconv.Itoa(result_sub)
	result_mul_ := strconv.Itoa(result_mul)
	result_div_ := strconv.Itoa(result_div)

	fmt.Printf(first + " + " + second + " = " + result_add_ + "\n")
	fmt.Printf(first + " - " + second + " = " + result_sub_ + "\n")
	fmt.Printf(first + " * " + second + " = " + result_mul_ + "\n")
	fmt.Printf(first + " / " + second + " = " + result_div_ + "\n")
}

func read() string {
	scanner := bufio.NewScanner(os.Stdin)
	var txt string
	if scanner.Scan() {
		txt = scanner.Text()
	}
	return txt
}

func to_int(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return integer
}
