package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	line := readFromFile()
	fmt.Printf("Value %v\n", firstInstruction(line, 12, 2))
	noun, verb := secondPart(line)
	fmt.Printf("Noun %v, %v : %v\n", noun, verb, (100*noun + verb))
}

func readFromFile() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Open file failed")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func firstInstruction(line string, noun int, verb int) int {
	instructions := line2instr(line)
	instructions[1] = noun
	instructions[2] = verb
	instructions = computeProgramm(instructions)
	return instructions[0]
}

func secondPart(line string) (int, int) {
	noun := 0
	verb := 0
Loop:
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			first := firstInstruction(line, noun, verb)
			if first == 19690720 {
				break Loop
			}
		}
	}
	return noun, verb
}

func line2instr(instr string) []int {
	line := fmt.Sprintf("[%v]", instr)
	var instructions []int
	err := json.Unmarshal([]byte(line), &instructions)
	if err != nil {
		log.Fatal(err)
	}
	return instructions
}

func computeProgramm(instr []int) []int {
	return next(0, instr)
}

func next(index int, instr []int) []int {
	fmt.Printf("Instruction %v à l'index %v\n", instr[index], index)
	switch instr[index] {
	case 1:
		sum := instr[instr[index+1]] + instr[instr[index+2]]
		instr = changeValue(instr, instr[index+3], sum)
		return next(index+4, instr)
	case 2:
		mult := instr[instr[index+1]] * instr[instr[index+2]]
		instr = changeValue(instr, instr[index+3], mult)
		return next(index+4, instr)
	default:
		return instr
	}
}

func changeValue(instr []int, index int, value int) []int {
	if index >= len(instr) {
		newInstr := make([]int, index+1, index+1)
		copy(newInstr, instr[:])
		instr = newInstr
	}
	instr[index] = value
	return instr
}
