package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var noun, verb, adj, adv string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a noun: ")
	if scanner.Scan() {
		noun = scanner.Text()
	}

	fmt.Print("Enter a verb: ")
	if scanner.Scan() {
		verb = scanner.Text()
	}

	fmt.Print("Enter an adjective: ")
	if scanner.Scan() {
		adj = scanner.Text()
	}

	fmt.Print("Enter an adverb: ")
	if scanner.Scan() {
		adv = scanner.Text()
	}

	fmt.Printf("Do you " + verb + " your " + adj + " " + noun + " " + adv + "? That's hilarious!\n")
}
