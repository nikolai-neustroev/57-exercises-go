package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var quote, author string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("What is the quote?")
	if scanner.Scan() {
		quote = scanner.Text()
	}

	fmt.Println("Who said it?")
	if scanner.Scan() {
		author = scanner.Text()
	}

	fmt.Printf(author + " says, \"" + quote + "\"\n")
}
