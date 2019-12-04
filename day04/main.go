package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func main() {
	input := "137683-596253"
	regex := regexp.MustCompile(`^(\d{6})-(\d{6})$`)
	match := regex.FindStringSubmatch(input)
	min, max := match[1], match[2]
	pass := passwords(min, max)
	fmt.Printf("NbPassword : %v\n", len(pass))
	fmt.Printf("NbPassword with strict double: %v\n", nbPasswordWithStrictDouble(pass))
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func nbPasswordWithStrictDouble(nums []int) int {
	res := 0
	for _, num := range nums {
		v := strconv.Itoa(num)
		ints := stringToInts(v)
		if hasStrictDouble(ints) {
			res++
		}
	}
	return res
}

func hasStrictDouble(num []int) bool {
	prev := -1
	for i, n := range num {
		if prev == n && ((i == len(num)-1) || num[i+1] != n) && (i < 2 || n != num[i-2]) {
			return true
		}
		prev = n
	}
	return false
}

func passwords(minS string, maxS string) []int {
	passwd := make([]int, 0, 1000)
	mi := toInt(minS)
	ma := toInt(maxS)
	for i := mi; i <= ma; i++ {
		if valid(i) {
			passwd = append(passwd, i)
		}
	}
	return passwd
}

func stringToInts(s string) []int {
	res := make([]int, len(s))
	for i, runeV := range s {
		res[i] = toInt(string(runeV))
	}
	return res
}

func valid(num int) bool {
	str := strconv.Itoa(num)
	ints := stringToInts(str)
	if !hasDouble(ints) {
		return false
	}
	if !isAsc(ints) {
		return false
	}
	return true
}

func isAsc(num []int) bool {
	prev := num[0]
	for _, v := range num {
		if prev > v {
			return false
		}
		prev = v
	}
	return true
}

func hasDouble(num []int) bool {
	prev := -1
	for _, v := range num {
		if prev == v {
			return true
		}
		prev = v
	}
	return false
}
