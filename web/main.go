package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var input Input
var entries []string

func main() {
	input := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Holds the string that scanned
		text := scanner.Text()
		text = strings.TrimSpace(text)
		if len(text) != 0 {
			fmt.Println(text)
			//exclude lines that are comments
			if strings.Contains(text, "#") {
				continue
			}
			input = append(input, text)
		} else {
			continue
		}
	}
	fmt.Println()
	if len(input) > 1 {
		doOperations(input)
	} else {
		fmt.Println("Not enough input to process. Please try again")
	}

}
