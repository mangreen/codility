# Reference
- https://leetcode.com/discuss/interview-question/3974163/CTC-2024-Intern-OA-Question-What-is-Optimized-Solution
- https://www.reddit.com/r/leetcode/comments/19fn39n/help_on_microsoft_oa_problem/

# TotalWaitingTime
Given a list of time estimates for several tasks, calculate the total waiting time for those tasks if they are processed cyclically in a round robin fashion (one time unit at a time).

## Task description
There are N clients who have ordered N handmade items. The K-th client ordered exactly one item that takes T[K] hours to make. There is only oneemployee who makes items for clients, and they work in the following manner:
- spend one hour making the first item;
- if the item is finished, the employee delivers it to the client immediately;
- if the item is not finished, they put it after the N-th item for further work;
- the employee starts making the next item.

For example for T = [3, 1, 2], the employee spends 6 hours making items in the following order: [1, 2, 3, 1, 3, 1]. The first client waited 6 hours for theiritem, the second client received their item after 2 hours and the third client after 5 hours. What is the total time that clients need to wait for allordered items? For the above example, the answer is 6 + 2 + 5 = 13.

As the result may be large, return its last nine digits without leading zeros (in other words, return the result modulo 10^9).

## Write a function:
    python: def solution(T: List[int]) -> int
    golang: func solution(T []int) int

that, given an array of integers T of length N, returns the total time that the clients need to wait (modulo 10^9).

## Examples:
1. Given T = [3, 1, 2], the function should return 13 as explained above.
2. Given T = [1, 2, 3, 4], the function should return 24. The employee prepares the items in the following order: 1, 2, 3, 4, 2, 3, 4, 3, 4, 4. The fi rst clientwaited for 1 hour, the second client for 5 hours, the third client for 8 hours, and the fourth client for 10 hours. The total waiting time of all clients is 1+ 5 + 8 + 10 = 24 hours.
3. Given T = [7, 7, 7], the function should return 60.
4. Given T = [10000], the function should return 10000.

## Write an efficient algorithm for the following assumptions:
- N is an integer within the range [1..100,000];
- each element of array T is an integer within the range [1..10,000].