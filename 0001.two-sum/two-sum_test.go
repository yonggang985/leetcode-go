package Problem0001

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	input  []int
	target int
}

type ans struct {
	output []int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)
	qs := []question{
		question{
			p: para{
				input:  []int{3, 2, 4},
				target: 6,
			},
			a: ans{
				output: []int{1, 2},
			},
		},
		question{
			p: para{
				input:  []int{3, 2, 4},
				target: 8,
			},
			a: ans{
				output: nil,
			},
		},
	}
	for _, v := range qs {
		a, p := v.a, v.p
		ast.Equal(a.output, twoSum(p.input, p.target), "input: %v", p)
	}
}
