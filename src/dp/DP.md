# Dynamic Programing

## Identifying DP problems:

- If you are given a problem, which can be broken down into smaller sub-problems, and these smaller sub-problems (Optimal substructure) can still be broken into smaller ones - and if you manage to find out that there are some overlapping subproblems, then you've encountered a DP problem.
    - Optimal Substructure
    - Overlapping sub-problem
  - Majority of the Dynamic Programming problems can be categorized into two types:
    - Optimization problems:
        - expect you to select a feasible solution, so that the value of the required function is minimized or maximized.
    - Combinatorial problems:
      - expect you to figure out the number of ways to do something, or the probability of some event happening.
  - Every Dynamic Programming problem has a schema to be followed:
    - Show that the problem can be broken down into optimal sub-structures.
    - Recursively define the value of the solution by expressing it in terms of optimal sub-structures.
    - Compute the value of the sub-structures in top-down (recursion) / Bottom-up (Iteration) fashion.
    - For overlapping subproblems use memoization to optimize the solution.
    - Construct an optimal solution from the computed information.

## Bottom Up
- I'm going to learn programming. Then, I will start practicing. Then, I will start taking part in contests. Then, I'll practice even more and try to improve. After working hard like crazy, I'll be an amazing coder.
- In Bottom Up, you start with the small solutions and then build up.
- Bottom-up is implemented by identifying the value of the base case or the optimal substructure and going up the ladder to find the solution to the larger problem.
- It is generally implemented in an iterative way and using a simple array.

## Top Down
- I will be an amazing coder. How? I will work hard like crazy. How? I'll practice more and try to improve. How? I'll start taking part in contests. Then? I'll be practicing. How? I'm going to learn programming.
- In Top Down, you start building the big solution right away by explaining how you build it from smaller solutions.
- In this approach we start from the larger problem and go down solving the smaller problem in a recursive manner.
- It is generally achieved in a recursive way.

## V.V.Important Rules:
- For problems where same question is asked with different values , the don’t go for space complexity O(1)
  - Identifying DP:
    - While solving a problem to identify subproblems, a strategy could be to
    - **Approach 1:**
      - Identify the Greedy approach for finding the suitable solution function for the problem
      - Then apply all possibilities to the solution function to see if the recursive state space tree has an overlapping subproblem and optimal substructure.

    - **Approach 2:**
      - Find the formula for solving the problem using Brute force
      - Then draw the recursive state space tree based on top down approach to see if there is an optimal substructure and overlapping subproblem or not.
      
    - **Approach 3:**
      - Use the Greedy approach to convert a complex problem into a simpler one and then use DP for the final solution.
      - Selection Rejection Approach to solving:
        - For identifying this approach there is generally a selection criteria in the question and that can tell us to use this approach.
        - To find the solution set and generate a recursion state space tree use, based on the question, a selection / rejection criteria for an element to draw the state space tree.
        - Find what will be the solution if A[i] is selected and what will be the solution if A[i] is not selected , that will create the state space tree.
        - Use the solution generally till A[i] , don’t consider future A[i]’s
    - Caution:
      - Don’t blindly think of substructure as mathematical induction where to find a solution for N we consider a solution to N-1 , N-2 etc..
      
## Final Solution for N:
- Basic strategy is to find all possible solutions, generate state space tree and derive the final solution.
- Now to solve for N , get all the solutions one step before N in the recursive state space tree and derive a result for N.
- Solution for N will be generally:
  - Dp[n] = some function of (Dp[n-x]), where x is the way in which the values one step before n are derived based on the identified DP approaches.
- Tips:
  - For the base case of Prefix sum always identify and define a base case to avoid confusions and conditional complex blocks.