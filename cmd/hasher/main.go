package main

import (
	"bufio"
	"fmt"
	"os"

	"edwingarcia.dev/github/jhunterscore/pkg/validation"
)

func main() {
	// Buffer to store the input.
	buff := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the string to hash:")

	str, _, err := buff.ReadLine()
	if err != nil {
		panic(err)
	}

	hashedStr, err := validation.HashPassword(str)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Original: %s\nHashed: %s\n", str, hashedStr)
}
