package dp

import "math"

/*
*
Problem Description
Given an integer A. Return minimum count of numbers, sum of whose squares is equal to A.
Count of numbers should be minimised which sums upto A and are perfect squares.

Problem Constraints
1 <= A <= 105

Input Format
First and only argument is an integer A.

Output Format
Return an integer denoting the minimum count.

Example Input
Input 1:

	A = 6

Input 2:

	A = 5

Example Output
Output 1:

	3

Output 2:

	2

Example Explanation
Explanation 1:

	Possible combinations are : (12 + 12 + 12 + 12 + 12 + 12) and (12 + 12 + 22).
	Minimum count of numbers, sum of whose squares is 6 is 3.

Explanation 2:

	We can represent 5 using only 2 numbers i.e. 12 + 22 = 5
*/

/*
*
Solutions:

 1. brute force or greedy solution is:
    a. find the largest perfect square less than or equal to A
    b. x = A - largest perfect square
    do above until x is 0
 2. greedy may fail as we are not sure if itis correct:
    - for each number between 2 to A
    a. find all perfect squares less than equal to A
    b. for each squares do:
    - dp[i] = min(i, dp[i - x^2]
*/
func MinSquares(A int) int {
	dp := make([]int, A+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= A; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(1+dp[i-(j*j)])))
		}
	}

	return dp[A]
}
