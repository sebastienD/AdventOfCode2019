package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start    point
	end      point
	vertical bool
}

func newLine(s point, e point) line {
	return line{s, e, s.x == e.x}
}

func (l line) cross(o line) (bool, point) {
	if l.vertical == o.vertical {
		return false, point{0, 0}
	}
	if l.vertical {
		return crossVertical(l, o)
	}
	return crossVertical(o, l)
}

func crossVertical(l line, hl line) (bool, point) {
	max := math.Max(float64(l.start.y), float64(l.end.y))
	min := math.Min(float64(l.start.y), float64(l.end.y))

	max2 := math.Max(float64(hl.end.x), float64(hl.start.x))
	min2 := math.Min(float64(hl.end.x), float64(hl.start.x))

	if min <= float64(hl.end.y) && max >= float64(hl.end.y) && min2 <= float64(l.end.x) && max2 >= float64(l.end.x) {
		return true, point{l.start.x, hl.end.y}
	}
	return false, point{0, 0}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	path1 := scanner.Text()
	scanner.Scan()
	path2 := scanner.Text()

	distance := dist(path1, path2)

	fmt.Printf("Dist is %v\n", distance)
}

func dist(p1 string, p2 string) int {
	instructions1 := toInstruction(p1)
	lines1 := instructions2lines(instructions1)

	instructions2 := toInstruction(p2)
	lines2 := instructions2lines(instructions2)

	points := cross(lines1, lines2)
	dist := nearFromOrigin(points)

	return dist
}

func nearFromOrigin(points []point) int {
	dist := float64(0)
	for _, p := range points {
		val := math.Abs(float64(p.x)) + math.Abs(float64(p.y))
		if val < dist || dist == 0 {
			dist = val
		}
	}
	return int(dist)
}

func cross(lines1 []line, lines2 []line) []point {
	points := make([]point, 0)
	for _, l1 := range lines1 {
		for _, l2 := range lines2 {
			res, point := l1.cross(l2)
			if res && (point.x != 0 || point.y != 0) {
				points = append(points, point)
			}
		}
	}
	return points
}

func toInstruction(path string) []string {
	return strings.Split(path, ",")
}

func instructions2lines(instrs []string) []line {
	lines := make([]line, 0)
	start := point{0, 0}
	for _, instr := range instrs {
		line := instr2line(start, instr)
		lines = append(lines, line)
		start = line.end
	}
	return lines
}

func instr2line(start point, instr string) line {
	move := instr[0:1]
	num := toi(instr[1:])
	switch move {
	case "U":
		return newLine(start, point{start.x, start.y + num})
	case "D":
		return newLine(start, point{start.x, start.y - num})
	case "R":
		return newLine(start, point{start.x + num, start.y})
	case "L":
		return newLine(start, point{start.x - num, start.y})
	}
	log.Fatalf("Move %v unknown", move)
	return newLine(point{0, 0}, point{0, 0})
}

func toi(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
