// a Go program which prompts the user to enter
// a floating point number and
// prints the integer
// which is a truncated version of
// the floating point number that was entered

// Truncation is the process of removing the digits
// to the right of the decimal place

/* 
Algorithm for the program
- Input - ask the user
- Conversion - truncate
- Output - print 
*/

package main

import (
	"fmt"
)

func main() {
	// Input
	var enter float64
	fmt.Println("Hey, please enter your best floating-point number: ")
	fmt.Scan(&enter)

	// Truncate - convert the number to an integer
	truncNo := int(enter)

	// Output - Print the result
	fmt.Println("So the Truncated integer is:", truncNo)
}