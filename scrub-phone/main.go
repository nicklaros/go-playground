package main

import (
	"fmt"
)

// scrubPhoneNumber scrubs the input phone number according to the specified pattern.
func scrubPhoneNumber(input string) string {
	length := len(input)

	// If the length is 4 or less, return the input as is.
	if length <= 4 {
		return input
	}

	var prefix string
	var maskStart int

	// Determine the prefix and where to start masking.
	if length >= 10 {
		// Preserve the first 3 characters.
		prefix = input[:3]
		maskStart = 3
	} else {
		prefix = ""
		maskStart = 0
	}

	// Last 4 digits are always visible.
	last4Start := length - 4

	// Edge case: If there's nothing to mask, return the input as is.
	if maskStart >= last4Start {
		return input
	}

	// Generate the masked part.
	maskedPart := ""
	for i := maskStart; i < last4Start; i++ {
		maskedPart += "*"
	}

	// Extract the last 4 digits.
	last4Digits := input[last4Start:]

	// Combine the prefix, masked part, and last 4 digits.
	return prefix + maskedPart + last4Digits
}

func main() {
	// Sample inputs.
	inputs := []string{
		"+6282245937412",
		"082245937412",
		"0333123654",
		"085324",
		"85324",
		"911",
	}

	// Process and print each input.
	for _, input := range inputs {
		output := scrubPhoneNumber(input)
		fmt.Printf("Input: %s, Output: %s\n", input, output)
	}
}
