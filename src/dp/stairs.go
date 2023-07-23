package dp

/**
Problem Description
You are climbing a staircase and it takes A steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Return the number of distinct ways modulo 1000000007



Problem Constraints
1 <= A <= 105



Input Format
The first and the only argument contains an integer A, the number of steps.



Output Format
Return an integer, representing the number of ways to reach the top.



Example Input
Input 1:

 A = 2
Input 2:

 A = 3


Example Output
Output 1:

 2
Output 2:

 3


Example Explanation
Explanation 1:

 Distinct ways to reach top: [1, 1], [2].
Explanation 2:

 Distinct ways to reach top: [1 1 1], [1 2], [2 1].
*/

/*
*
Solution:
n = 0:

	--> 1 ways

n = 1:

	--> 0 -> 1 ; 1 ways

n = 2:

	2 ways: numofways(1) + numofways(0)

n =3:

	3 ways: numofways(2) + numofways(1)

Therefore: f(n) = f(n-1) + f(n-2)
*/
func ClimbStairs(A int) int {
	dp := make([]int, A+1)
	dp[0] = 1
	dp[1] = 1

	if A == 1 {
		return dp[1]
	}
	for i := 2; i <= A; i++ {
		//The sum can become larger than what int can hold and circle back to negetive,
		//SO modulus will keep the result of sum within int range
		dp[i] = (dp[i-2] + dp[i-1]) % 1000000007
	}
	return dp[A]
}
