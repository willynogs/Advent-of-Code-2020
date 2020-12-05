package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strings"
)

func calculateSeatID(input string) int {
	instructions := strings.Split(input, "")
	rowInstructions := instructions[:7]
	colInstructions := instructions[7:]
	rowMin := 0.0
	rowMax := 127.0
	colMin := 0.0
	colMax := 7.0
	for _, ins := range rowInstructions {
		op := (rowMin + rowMax) / 2
		if ins == "F" {
			rowMax = math.Floor(op)
		} else {
			rowMin = math.Ceil(op)
		}
	}
	for _, ins := range colInstructions {
		op := (colMin + colMax) / 2
		if ins == "L" {
			colMax = math.Floor(op)
		} else {
			colMin = math.Ceil(op)
		}
	}
	return int((rowMin * 8) + colMin)
}

func partOne(input []string) {
	highestID := -1
	for _, entry := range input {
		id := calculateSeatID(entry)
		if id > highestID {
			highestID = id
		}
	}
	fmt.Printf("Part 1 Answer: %d", highestID)
	fmt.Println()
}

func partTwo(input []string) {
	var seatIDs []int
	for _, entry := range input {
		id := calculateSeatID(entry)
		seatIDs = append(seatIDs, id)
	}
	sort.Ints(seatIDs)
	for i := 0; i < (len(seatIDs) - 2); i++ {
		low := seatIDs[i]
		mid := seatIDs[i+1]
		high := seatIDs[i+2]
		if !(low+1 == mid && mid+1 == high) {
			fmt.Printf("Part 2 Answer: %d", mid+1)
			fmt.Println()
			break
		}
	}
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	entries := strings.Split(string(content), "\n")
	partOne(entries)
	partTwo(entries)
}
