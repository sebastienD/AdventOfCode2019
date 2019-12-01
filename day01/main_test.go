package main

import (
	"fmt"
	"testing"
)

func TestShouldComputeFuelNeeded(t *testing.T) {
	var massData = []struct {
		mass         int
		fuelExpected int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
		{3, 0},
	}
	for _, data := range massData {
		fuelActual := computeFuel(data.mass)
		if fuelActual != data.fuelExpected {
			t.Errorf("Line:[%v] expected %v, actual %v\n", data.mass, data.fuelExpected, fuelActual)
		} else {
			fmt.Printf("Line:[%v] ok %v\n", data.mass, fuelActual)
		}
	}

}
