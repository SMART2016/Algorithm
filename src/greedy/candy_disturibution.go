package greedy

import (
	"fmt"
	"math"
)

/**
Problem Description
	N children are standing in a line. Each child is assigned a rating value.

	You are giving candies to these children subjected to the following requirements:

	Each child must have at least one candy.
	Children with a higher rating get more candies than their neighbors.
	What is the minimum number of candies you must give?

Problem Constraints
	1 <= N <= 105
	-109 <= A[i] <= 109

Input Format
	The first and only argument is an integer array A representing the rating of children.

Output Format
	Return an integer representing the minimum candies to be given.

Example Input
	Input 1:
	 A = [1, 2]

	Input 2:
	 A = [1, 5, 2, 1]


Example Output
	Output 1:
	 3

	Output 2:
	 7

Example Explanation
	Explanation 1:
		 The candidate with 1 rating gets 1 candy and candidate with rating 2 cannot get 1 candy as 1 is its neighbor.
		 So rating 2 candidate gets 2 candies. In total, 2 + 1 = 3 candies need to be given out.

	Explanation 2:
	 Candies given = [1, 3, 2, 1]
*/

/*
*
Solution:

	There are two condition to be met for assigning candy to a student optimally:
	If the element left of A[i] less that A[i]
	If the element right of A[i] is less than A[i]
	- So we will do this in 2 passes:
	- 1st Pass Left --> Right iteration: we check if the current element is greater than the element at the left , we assign the value of the current element
			as c[i-1] + 1 otherwise we assign it 1
	- 2nd Pass Right --> Left iteration: we check if A[i] > A[i+1], then we assign c[i] = max(c[i],c[i+1]+1)
*/
func DistCandy(A []int) int {
	c := make([]int, len(A))
	c[0] = 1
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			//c[i] = c[i-1] + 1
			c[i] = c[i-1] + 1
		} else {
			c[i] = 1
		}
	}

	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			c[i] = int(math.Max(float64(c[i]), float64(c[i+1]+1)))
		}
	}

	//Find sum of the elements in c
	sum := 0
	for i, _ := range c {
		sum = sum + c[i]
	}
	fmt.Println("candy := ", c)
	return sum
}
