package total_waiting_time

import (
	"math"
)

func solution(T []int) int {
	MOD := int(math.Pow(10, 9))

	ans := 0
    cnt := 0
    
    q := []int{}
    for i := range T {
        q = append(q, i)
    }
    
    for len(q)>0 {
        for _, idx := range q {
            cnt++
            
            q = q[1:]
            T[idx]--
            
            if T[idx] > 0 {
                q = append(q, idx)
            } else {
                ans += cnt
            }
        }
    }
    
    return ans % MOD
}