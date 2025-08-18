package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		clean := cleanInput(userInput)
		fmt.Printf("Your command was: %v\n", clean[0])
	}

}

func cleanInput(text string) []string {
	result := strings.ToLower(text)
	return strings.Fields(result)
}
