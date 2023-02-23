package main

import (
	"fmt"
	"hackerrank/golang"
)

func main() {
	// Solve Me First
	fmt.Println(golang.SolveMeFirst(30, 40))

	// Simple Array Sum
	simpleArraySum := []int32{2, 3, 4, 2, 5, 12, 20}
	fmt.Println(golang.SimpleArraySum(simpleArraySum))

	// Compare The Triplets
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}
	fmt.Println(golang.CompareTriplets(a, b))
}
