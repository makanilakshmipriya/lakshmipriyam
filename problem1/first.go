package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// removes hyphens from the given string
func removeHyphens(s string) string {
	return regexp.MustCompile("-").ReplaceAllString(s, "")
}

// checks if the given card number is valid (isValidCreditCardNumber)
func isValidCreditCardNumber(cardNumbr string) bool {
	// Defines regex pattern to match valid credit card numbers
	pattern := `^(?:4|5|6)\d{3}(-?\d{4}){3}$`

	// Compile the regex pattern
	regex := regexp.MustCompile(pattern)

	// Checkes if the card number matches the pattern
	if !regex.MatchString(cardNumbr) {
		return false
	}

	// Removes the hyphens from the card number
	cleanedCardNumb := removeHyphens(cardNumbr)

	// Check for consecutive repeated digits
	for i := 0; i < len(cleanedCardNumb)-3; i++ {
		if cleanedCardNumb[i] == cleanedCardNumb[i+1] && cleanedCardNumb[i] == cleanedCardNumb[i+2] && cleanedCardNumb[i] == cleanedCardNumb[i+3] {
			return false
		}
	}

	return true
}

func main() {
	// Scanner to read user input from terminal
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter credit card number (or type 'exit' to quit):")

	for scanner.Scan() {
		input := scanner.Text()
		// Exit loop if user types 'exit'
		if input == "exit" {
			break
		}

		// Check the validity of the card number
		isValid := isValidCreditCardNumber(input)
		validity := "Invalid"
		if isValid {
			validity = "Valid"
		}
		fmt.Printf("Credit Card Number: %s - %s\n", input, validity)
		fmt.Println("Enter another credit card number (or type 'exit' to quit):")
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
	}
}
