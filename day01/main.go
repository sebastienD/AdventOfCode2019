package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func computeFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func computeRecursiveFuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel + computeRecursiveFuel(fuel)
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
		total += computeFuel(toi(mass))
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
		total += computeRecursiveFuel(toi(mass))
	}

	fmt.Printf("Total %v\n", total)
}
