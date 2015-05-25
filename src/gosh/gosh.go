package main

import (
	"bufio"
	"fmt"
	"os"
	"parser"
)

func main() {
	var input string

	consoleReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, _ = consoleReader.ReadString('\n')

		tokens := parser.Tokenize(input)
		fmt.Println(tokens)

	}
}
