package main

//“Find length of longest substring of a given string of digits, such that sum of digits in the first half and second half of the substring is same. For example,
//Input: “142124”
//Output: 6
//The whole string is answer, because, sum of first 3 digits = sum of last 3 digits (1+4+2 = 1+2+4).”

func findSubsetwithEqlSumHalf(arr []int, length int) int {
	sum := init2DSlice(length)
	maxlength := 0
	for i := 0; i < len(arr); i++ {
		sum[i][i] = arr[i]

	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < i+(len(arr)/2); j++ {
			if j < len(arr) {
				sum[i][j] = sum[i][j-1] + arr[j]
				if length, ok := checkIsMaxLengthSubset(sum, i, j); ok {
					maxlength = length
				}
			}
		}
	}
	return maxlength
}

func init2DSlice(length int) [][]int {
	sum := make([][]int, length) // initialize a slice of dy slices
	for i := 0; i < length; i++ {
		sum[i] = make([]int, length) // initialize a slice of dx unit8 in each of dy slices
	}
	return sum
}

func checkIsMaxLengthSubset(sum [][]int, i int, j int) (int, bool) {
	len := j - i + 1
	if i-len >= 0 && j-len >= 0 && j-len != i {
		return len, sum[i-len][j-len] == sum[i][j]
	}
	return len, false
}
