package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgramm(t *testing.T) {
	var data = []struct {
		instructions  string
		instrExpected string
	}{
		{"1,0,0,0,99", "2,0,0,0,99"},
		{"2,3,0,3,99", "2,3,0,6,99"},
		{"2,4,4,5,99", "2,4,4,5,99,9801"},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
	}
	for _, d := range data {
		instructions := line2instr(d.instructions)
		actual := computeProgramm(instructions)
		fmt.Printf("%v\n", strings.Join(strings.Fields(fmt.Sprint(actual)), ","))
		instrActual := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(actual)), ","), "[]")
		assert.Equal(t, d.instrExpected, instrActual, "they should be equal")
	}
}
