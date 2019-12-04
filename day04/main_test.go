package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrictDouble(t *testing.T) {
	var data = []struct {
		num      []int
		expected bool
	}{
		{[]int{5, 8, 8, 8, 8, 8}, false},
		{[]int{5, 8, 8, 8, 8, 9}, false},
		{[]int{5, 8, 8, 8, 9, 9}, true},
		{[]int{5, 8, 8, 9, 9, 9}, true},
		{[]int{5, 9, 9, 9, 9, 9}, false},
	}
	for _, d := range data {
		actual := hasStrictDouble(d.num)
		assert.Equal(t, d.expected, actual, "Should be true", d.num)
	}
}
