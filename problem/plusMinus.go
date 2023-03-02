package problem

import "fmt"

func PlusMinus(arr []int32) {
	var pos, neg, zero float32
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			pos += 1
		}
		if arr[i] < 0 {
			neg += 1
		}
		if arr[i] == 0 {
			zero += 1
		}
	}
	// fmt.Println(pos,neg,zero)
	pos = pos / float32(len(arr))
	neg = neg / float32(len(arr))
	zero = zero / float32(len(arr))
	fmt.Printf("%f\n", pos)
	fmt.Printf("%f\n", neg)
	fmt.Printf("%f", zero)

}
