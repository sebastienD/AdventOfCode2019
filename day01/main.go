package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func computeFuel(mass int, rec bool) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	if rec {
		return fuel + computeFuel(fuel, rec)
	}
	return fuel
}

func toi(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func main() {
	filename := "input.txt"
	firstPart(filename)
	secondPart(filename)
}

func firstPart(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass := scanner.Text()
		total += computeFuel(toi(mass), false)
	}

	fmt.Printf("Total %v\n", total)
}

func secondPart(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass := scanner.Text()
		total += computeFuel(toi(mass), true)
	}

	fmt.Printf("Total %v\n", total)
}
