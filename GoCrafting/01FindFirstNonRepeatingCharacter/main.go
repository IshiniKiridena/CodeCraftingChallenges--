package main

import (
	"fmt"
	"os"
)

func FindFirstNonRepCharacter(inputStr string) byte {

	//frequency map to keep track of the character occurrences
	frequencyMap := make(map[byte]int)

	//loop through the string
	for i := 0; i < len(inputStr); i++ {
		// Check if the character is already in the map
		if _, exists := frequencyMap[inputStr[i]]; exists {
			//update value
			frequencyMap[inputStr[i]]++
		} else {
			//new entry
			frequencyMap[inputStr[i]] = 1
		}
	}

	// Loop through the string again to find the first non-repeating character
	for i := 0; i < len(inputStr); i++ {
		if frequencyMap[inputStr[i]] == 1 {
			return inputStr[i]
		}
	}
	return 0
}

func main() {
	// Check if at least one command-line argument (excluding the program name) is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your_program.go <input_string>")
		return
	}
	// Get the input string from the command-line arguments
	inputString := os.Args[1]
	byteValue := FindFirstNonRepCharacter(inputString)
	if byteValue == 0 {
		fmt.Println("Output : ", byteValue)
	} else {
		fmt.Println("Output : ", string(byteValue))
	}

	//! To run the code snippet the command -> go run main.go "<your string>"
}
