package arrSum

/*
*
returns sum of all integers in array
*/
func Sum(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	var sums []int
	for i := range numsToSum {
		sums = append(sums, Sum(numsToSum[i]))
	}
	return sums
}

// For every subarray, calculate sum of all elements except first, return resulting array
func SumAllTails(arrays ...[]int) (tailSumArr []int) {
	for _, currArr := range arrays {
		if len(currArr) == 0 {
			tailSumArr = append(tailSumArr, 0)
		} else {
			tailSumArr = append(tailSumArr, Sum(currArr[1:]))
		}
	}
	return tailSumArr
}
