package count_bananas

func solution(S string) int {
	cntB := 0
    cntA := 0
    cntN := 0
    
    for _, r := range S {
        switch r {
            case 'B':
                cntB++
            case 'A':
                cntA++
            case 'N':
                cntN++
        }
    }
    
    return min(cntB, cntA/3, cntN/2)
}
