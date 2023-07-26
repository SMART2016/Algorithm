package dp

import "math"

/*
Problem Description
Given a 2 x N grid of integers, A, your task is to choose numbers from the grid such that sum of these numbers is maximized.
However, you cannot choose two numbers that are adjacent horizontally, vertically, or diagonally.

Return the maximum possible sum.

Note: You are allowed to choose more than 2 numbers from the grid.


Problem Constraints
1 <= N <= 20000
1 <= A[i] <= 2000


Input Format
The first and the only argument of input contains a 2d matrix, A.


Output Format
Return an integer, representing the maximum possible sum.


Example Input
Input 1:

 A = [
        [1]
        [2]
     ]
Input 2:

 A = [
        [1, 2, 3, 4]
        [2, 3, 4, 5]
     ]


Example Output
Output 1:

 2
Output 2:

 8


Example Explanation
Explanation 1:

 We will choose 2 (From 2nd row 1st column).
Explanation 2:

 We will choose 3 (From 2nd row 2nd column) and 5 (From 2nd row 4th column).
*/

/*
Approach:
  - In this problem we can see that we can select ith element only from 1 row,if we select i from 1st row then we don't
    select i from second row and vice versa.
  - It makes sense to choose the larger number from between the column x of row 1 and 2, as we need to maximize the sum.
  - element = max(A[0][i],A[1][i])
  - This converts the problem to an 1d array logically
  - Now we apply select/reject strategy on the result 1d array to reach to the final solution.
  - If A[i] selected , then maxsum = A[i] + sum[i-2]
  - if A[i] is rejected , then maxsum = sum[i-1]
  - therefore the maxsum till ith element = max(A[i] + sum[i-2] , sum[i-1])
*/
func Adjacent(A [][]int) int {
	numOfCols := len(A[0])
	maxSum := 0
	sumprev := 0
	sumprevtoprev := 0
	for i := 0; i < numOfCols; i++ {
		if i > 0 {
			maxSum = max(max(A[0][i], A[1][i])+sumprevtoprev, sumprev)
		} else {
			maxSum = max(max(A[0][i], A[1][i])+sumprevtoprev, sumprev)
		}
		sumprevtoprev = sumprev
		sumprev = maxSum
	}
	return maxSum
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}
