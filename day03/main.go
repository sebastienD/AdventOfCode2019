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
	start     point
	end       point
	vertical  bool
	direction string
}

func newLine(s point, e point, d string) line {
	return line{s, e, s.x == e.x, d}
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

	nbSteps := lessStep(path1, path2)
	fmt.Printf("NbSteps is %v\n", nbSteps)
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
		return newLine(start, point{start.x, start.y + num}, move)
	case "D":
		return newLine(start, point{start.x, start.y - num}, move)
	case "R":
		return newLine(start, point{start.x + num, start.y}, move)
	case "L":
		return newLine(start, point{start.x - num, start.y}, move)
	}
	log.Fatalf("Move %v unknown", move)
	return newLine(point{0, 0}, point{0, 0}, "")
}

func toi(in string) int {
	num, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func lessStep(p1 string, p2 string) int {
	instructions1 := toInstruction(p1)
	lines1 := instructions2lines(instructions1)

	instructions2 := toInstruction(p2)
	lines2 := instructions2lines(instructions2)

	points := cross(lines1, lines2)

	nbSteps1, crossPoint := firstCross(lines1, points)
	fmt.Printf("First Cross is %v in nbSteps %v\n", crossPoint, nbSteps1)

	nbSteps2 := nbStepsTo(crossPoint, lines2)

	return nbSteps1 + nbSteps2
}

func firstCross(lines []line, crosses []point) (int, point) {
	nbSteps := 0
	for _, l := range lines {
		for _, p := range crosses {
			ok, nb := pointIsIn(l, p)
			if ok {
				return nbSteps + nb, p
			}
		}
		nbSteps += nbStepsIn(l)
	}
	return 0, point{0, 0}
}

func pointIsIn(l line, p point) (bool, int) {
	if l.vertical {
		max := math.Max(float64(l.start.y), float64(l.end.y))
		min := math.Min(float64(l.start.y), float64(l.end.y))
		if p.x == l.start.x && float64(p.y) >= min && float64(p.y) <= max {
			res := true
			if l.direction == "U" {
				return res, int(math.Abs(float64(p.y - l.start.y)))
			}
			return res, int(math.Abs(float64(l.start.y - p.y)))
		}
		return false, 0
	}

	max := math.Max(float64(l.start.x), float64(l.end.x))
	min := math.Min(float64(l.start.x), float64(l.end.x))
	if p.y == l.start.y && float64(p.x) >= min && float64(p.x) <= max {
		res := true
		if l.direction == "R" {
			return res, int(math.Abs(float64(p.x - l.start.x)))
		}
		return res, int(math.Abs(float64(l.start.x - p.x)))
	}
	return false, 0
}

func nbStepsIn(l line) int {
	if l.vertical {
		return int(math.Abs(float64(l.end.y - l.start.y)))
	}
	return int(math.Abs(float64(l.end.x - l.start.x)))
}

func nbStepsTo(p point, lines []line) int {
	nbSteps := 0
	for _, l := range lines {
		ok, nb := pointIsIn(l, p)
		if ok {
			return nbSteps + nb
		}
		nbSteps += nbStepsIn(l)
	}
	return 0
}
