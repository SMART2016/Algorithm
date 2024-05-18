package main

func countSquareSubMatrix(matrix [][]int) int {
	dp := make([][]int, len(matrix)+1)
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0])+1)
		dp[i][0] = matrix[i][0]
	}

	for j := 0; j < len(matrix[0]); j++ {
		dp[0][j] = matrix[0][j]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
			}
		}
	}

	totalSquareSubMatrix := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			totalSquareSubMatrix += dp[i][j]
		}
	}

	return totalSquareSubMatrix
}

func TestCountSquareSubMatrix() int {
	matrix := [][]int{{1, 0, 1}, {1, 1, 0}, {1, 1, 0}}
	return countSquareSubMatrix(matrix)
}
