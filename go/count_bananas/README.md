# Reference
- https://www.chegg.com/homework-help/questions-and-answers/1-countbananas-calculate-many-times-print-word-banana-using-letters-given-string-s-task-de-q113659844

# CountBanbanas
Calculate how many times you can print the word "BANANA" using the letters given in string S.

## Task description
A string S made of uppercase English letters is given. In one move, six letters forming the word "BANANA" (1 'B', 3 'A' and 2 'N') can be deleted from S. What is the max number times such a move can be applied to S? 

## Write a function: 
    java: class solution { public int solution(Strine S); } 
    c++: int solution(string &S)
    golang func solution(S string) int

that, given a string S of length N, returns the maximum number of moves that can be applied. 

## Examples: 
1. Given S= = NAABXXAN', the function should retum 1. 
2. Given S = 'NAANAAXNABABYNNBZ', the function should return 2. 
3. Given S= = QABAAAWOBL, the function should retum 0 . 

## Write an efficient algorithm for the following assumptions: 
- N is an integer within the range [1..100,000]; 
- string S is made only of uppercase letters (A−z).