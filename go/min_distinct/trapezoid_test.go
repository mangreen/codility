package min_distinct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_121(t *testing.T) {
	assert := assert.New(t)

	ans := solution([]int{1, 2, 1})

	assert.Equal(2, ans, "2, Increase 2 times [1, 2, 1] -> [1, 2, 3]")
}

func Test_solution_2144(t *testing.T) {
	assert := assert.New(t)

	ans := solution([]int{2, 1, 4, 4})

	assert.Equal(1, ans, "1, Decrease 1 times [2, 1, 4, 4] -> [2, 1, 3, 4]")
}

func Test_solution_623563(t *testing.T) {
	assert := assert.New(t)

	ans := solution([]int{6, 2, 3, 5, 6, 3})

	assert.Equal(4, ans, "4, Decrease 4 times [6, 2 , 3, 5, 6, 3] -> [6, 2 , 1, 5, 4, 3]")
}

func Test_solution_222555888(t *testing.T) {
	assert := assert.New(t)

	ans := solution([]int{2, 2, 2, 5, 5, 5, 8, 8, 8})

	assert.Equal(6, ans, "6, Decrease 4 times [2, 2, 2, 5, 5, 5, 8, 8, 8] -> [1, 2, 3, 4, 5, 6, 7, 8, 9]")
}