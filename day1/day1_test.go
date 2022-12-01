package day1

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const input = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestPart1Example(t *testing.T) {
	assert.Equal(t, 24000, part1(bufio.NewScanner(strings.NewReader(input))))
}
