// Algorithm

// Input
// Search
// Check Conditions
// Output

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input
	var userInput string
	fmt.Print("Heyo, enter a string\n")
	fmt.Scan(&userInput)

	// Search and Check Conditions
	if strings.HasPrefix(strings.ToLower(userInput), "i") && strings.HasSuffix(strings.ToLower(userInput), "n") && strings.Contains(strings.ToLower(userInput), "a") {
		// Output
		fmt.Println("Found!")
	} else {
		// Output
		fmt.Println("Not Found!")
	}
}