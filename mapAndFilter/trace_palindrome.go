package main

import "fmt"

// Tracing version of IsPalindrome to show step-by-step execution
func IsPalindromeTrace(x int) bool {
	fmt.Printf("\n=== Tracing IsPalindrome(%d) ===\n", x)
	
	// Handle negative numbers by checking their absolute value
	if x < 0 {
		fmt.Printf("  Step 1: %d is negative, converting to absolute value\n", x)
		x = -x // Convert to positive
		fmt.Printf("  Step 1: x is now %d\n", x)
	} else {
		fmt.Printf("  Step 1: %d is not negative, proceeding\n", x)
	}
	
	// Handle single digits (including 0)
	if x < 10 {
		fmt.Printf("  Step 2: %d is a single digit, returning true\n", x)
		return true
	}
	
	fmt.Printf("  Step 2: %d is not a single digit, proceeding to reverse\n", x)
	
	num := x
	reversed := 0
	step := 3
	
	fmt.Printf("  Step %d: Starting reversal process\n", step)
	fmt.Printf("  Step %d: num = %d, reversed = %d\n", step, num, reversed)
	
	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		fmt.Printf("  Step %d: digit = %d, reversed = %d*10 + %d = %d\n", step, digit, reversed/10, digit, reversed)
		num /= 10
		fmt.Printf("  Step %d: num = %d / 10 = %d\n", step, num*10, num)
		step++
	}
	
	result := x == reversed
	fmt.Printf("  Final: x = %d, reversed = %d, x == reversed = %t\n", x, reversed, result)
	fmt.Printf("=== Result: %t ===\n", result)
	
	return result
}

func main() {
	fmt.Println("TRACING IsPalindrome FUNCTION")
	fmt.Println("=============================")
	
	// Test cases to trace
	testCases := []int{
		121,    // Positive palindrome
		-121,   // Negative palindrome
		123,    // Positive non-palindrome
		-123,   // Negative non-palindrome
		0,      // Zero
		1,      // Single digit
		10,     // Ends with zero
		1221,   // Even length palindrome
		12321,  // Odd length palindrome
	}
	
	for _, num := range testCases {
		IsPalindromeTrace(num)
		fmt.Println()
	}
	
	// Also test the original function to compare
	fmt.Println("COMPARISON WITH ORIGINAL FUNCTION:")
	fmt.Println("==================================")
	for _, num := range testCases {
		original := IsPalindrome(num)
		fmt.Printf("IsPalindrome(%d) = %t\n", num, original)
	}
} 