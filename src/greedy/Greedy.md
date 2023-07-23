# Greedy
## Optimization problems:
expect you to select a feasible solution, so that the value of the required function is minimized or maximized.

## Strategies:
- Correcting an incorrect step from the past.
- Try best possibility via induction
  - Solve for N = 1
  - Assume the same for N = N - 1
    - In this step we try to identify what we can do for N-1 problems , which will help me reach , maximize the possibility of solving the Nth sub problem.
  - Solve for N = N
  - If there are multiple rules for reaching the solution , then handle one rule at a time with multiple passes.
  - Try small inputs in increasing order to find a pattern.

## Spotting a Greedy Algorithm
  - Spotting a greedy algorithm problem can be important in algorithmic problem-solving and competitive programming. Greedy algorithms are a specific type of algorithmic paradigm that make locally optimal choices at each step with the hope of finding a global optimum. Here are some characteristics that can help you identify a greedy algorithm problem:
    - Optimization problem with a natural greedy choice: Greedy algorithms are commonly used to solve optimization problems where you want to maximize or minimize a certain objective function. In these problems, there is a natural, intuitive choice at each step that seems to be the best option to reach the optimal solution.
    Greedy choice property: The greedy choice property means that at each step of the algorithm, the choice made should lead to a locally optimal solution. In other words, the choice made at a particular step should be the best possible choice at that moment without considering the overall future consequences.
    Optimal substructure: Greedy algorithms often have the optimal substructure property, which means that the optimal solution to the problem can be constructed from the optimal solutions of its subproblems.
    Proof of correctness by greedy choice property: The correctness of a greedy algorithm is usually proved by showing that the locally optimal choices made at each step eventually lead to the global optimal solution.
    No backtracking: Greedy algorithms make decisions at each step and do not backtrack to undo choices. Once a choice is made, it's final.
    Lack of overlapping subproblems: Unlike dynamic programming, which typically has overlapping subproblems, greedy algorithms do not encounter overlapping subproblems because they make decisions without revisiting the same subproblems.
    
  - Examples of greedy algorithm problems:
    - Activity Selection Problem: Given a set of activities with start and finish times, find the maximum number of activities that can be performed, assuming only one activity can be performed at a time.
    - Huffman Coding: Given a set of characters and their frequencies, construct a binary tree such that characters with higher frequencies have shorter codes.
    - Coin Change Problem: Given a set of coin denominations and a target amount, find the minimum number of coins needed to make the change.
    - Fractional Knapsack Problem: Given a set of items with weights and values, fill a knapsack with a maximum weight capacity to maximize the total value of the items included.
  - Remember that not all optimization problems can be solved using a greedy approach. It's essential to analyze the problem carefully and identify whether the greedy approach is appropriate. Sometimes, greedy algorithms might not provide the optimal solution for the given problem, in which case, other paradigms like dynamic programming or divide and conquer should be considered.
