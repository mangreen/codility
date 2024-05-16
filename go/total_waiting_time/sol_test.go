package total_waiting_time

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_312(t *testing.T) {
	assert := assert.New(t)

	ans := solution([]int{3, 1, 2}) // 1,2,3,1,3,1

	assert.Equal(13, ans, "13, 6 + 2 + 5 = 13")
}

func Test_solution_1234(t *testing.T) {
	assert := assert.New(t)

	ans := solution([]int{1, 2, 3, 4}) // 1,2,3,4,2,3,4,3,4,4

	assert.Equal(24, ans, "24, 1 + 5 + 8 + 10 = 24")
}

func Test_solution_777(t *testing.T) {
	assert := assert.New(t)

	ans := solution([]int{7, 7, 7}) // 1,2,3,1,2,3,...,1,2,3

	assert.Equal(60, ans, "60, 19 + 20 + 21 = 60")
}

func Test_solution_1000(t *testing.T) {
	assert := assert.New(t)

	ans := solution([]int{1000}) // 1,1,1,1,...,1

	assert.Equal(1000, ans, "1000")
}