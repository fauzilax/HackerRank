package problem

func CompareTriplets(a []int32, b []int32) []int32 {
	var aScore, bScore int32
	var result []int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			aScore += 1
		}
		if a[i] < b[i] {
			bScore += 1
		}
		if a[i] == b[i] {
			aScore += 0
			bScore += 0
		}
	}
	result = append(result, aScore)
	result = append(result, bScore)

	return result
}
