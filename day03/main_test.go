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
		{newLine(point{1, 2}, point{1, 8}), newLine(point{0, 4}, point{8, 4}), point{1, 4}, true},
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
