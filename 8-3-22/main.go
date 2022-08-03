package main

import "fmt"

// Create a list of numbers
var numList = []int{1, 2, 3, 4, 5, 6}

// Set k equal to a number
var k int = 6

func main() {
	// Iterate over the list twice
	for _, i := range numList {
		for _, j := range numList {
			z := add(i, j)
			comp := compare(z, k)
			// Only print those that do match
			if comp {
				fmt.Printf("%v + %v = %v\n", i, j, z)
				fmt.Printf("Are %v and %v equal? - %v\n", z, k, comp)
			}
		}
	}
}

// Add two values
func add(x int, y int) int {
	var sum int = x + y
	return sum
}

// Compare two values
func compare(x int, y int) bool {
	if x == y {
		return true
	}
	return false
}
