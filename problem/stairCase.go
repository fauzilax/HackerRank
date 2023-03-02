package problem

import "fmt"

func StairCase(n int32) {
	// Write your code here
	var i, j, k int32
	for i = 0; i < n; i++ {
		for j = i; j < n-1; j++ {
			fmt.Print(" ")
		}
		for k = 0; k < i+1; k++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}

}
