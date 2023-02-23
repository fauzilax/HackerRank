package problem

func SimpleArraySum(ar []int32) int32 {
	var sum int32
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum

}
