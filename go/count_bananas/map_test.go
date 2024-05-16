package count_bananas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution_NAABXXAN(t *testing.T) {
	assert := assert.New(t)

	ans := solution("NAABXXAN")

	assert.Equal(1, ans, "1, 'NAABXXAN' -> 'XX'")
}

func Test_solution_NAANAAXNABABYNNBZ(t *testing.T) {
	assert := assert.New(t)

	ans := solution("NAANAAXNABABYNNBZ")

	assert.Equal(2, ans, "2, 'NAANAAXNABABYNNBZ' -> 'NAAXNABYNBZ' -> 'XBYNZ'")
}

func Test_solution_QABAAAWOBL(t *testing.T) {
	assert := assert.New(t)

	ans := solution("QABAAAWOBL")

	assert.Equal(0, ans, "0, 'QABAAAWOBL'")
}