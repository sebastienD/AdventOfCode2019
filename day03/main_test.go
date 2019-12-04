package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	var paths = []struct {
		path1        string
		path2        string
		distExpected int
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 6},
	}
	for _, p := range paths {
		actualDist := dist(p.path1, p.path2)
		assert.Equal(t, p.distExpected, actualDist, "should be equals")
	}
}

func TestCross(t *testing.T) {
	var data = []struct {
		line1           line
		line2           line
		pointExpected   point
		interscExpected bool
	}{
		{newLine(point{1, 2}, point{1, 8}, "R"), newLine(point{0, 4}, point{8, 4}, "L"), point{1, 4}, true},
	}
	for _, d := range data {
		ok, p := d.line1.cross(d.line2)
		assert.Equal(t, d.interscExpected, ok, "should be equals")
		assert.Equal(t, d.pointExpected, p, "should be equals")
		ok, p = d.line2.cross(d.line1)
		assert.Equal(t, d.interscExpected, ok, "should be equals")
		assert.Equal(t, d.pointExpected, p, "should be equals")
	}
}

func TestSteps(t *testing.T) {
	var data = []struct {
		path1           string
		path2           string
		nbStepsExpected int
	}{
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 610},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
	}
	for _, d := range data {
		nbSteps := lessStep(d.path1, d.path2)
		assert.Equal(t, d.nbStepsExpected, nbSteps, "Should be equals")
	}
}
