package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func partOne(input []string) {
	treesHit := 0

	for i, row := range input {
		if string(row[(i*3)%len(row)]) == "#" {
			treesHit++
		}
	}

	fmt.Printf("Part 1 Answer: %d", treesHit)
	fmt.Println()
}

func partTwo(input []string) {
	slopeOneHits := 0
	slopeTwoHits := 0
	slopeThreeHits := 0
	slopeFourHits := 0
	slopeFiveHits := 0

	for i, row := range input {
		if string(row[i%len(row)]) == "#" {
			slopeOneHits++
		}

		if string(row[(i*3)%len(row)]) == "#" {
			slopeTwoHits++
		}

		if string(row[(i*5)%len(row)]) == "#" {
			slopeThreeHits++
		}

		if string(row[(i*7)%len(row)]) == "#" {
			slopeFourHits++
		}

		if i%2 != 0 && string(row[(i*2)%len(row)]) == "#" {
			slopeFiveHits++
		}
	}

	product := slopeOneHits * slopeTwoHits * slopeThreeHits * slopeFourHits * slopeFiveHits
	fmt.Printf("Part 2 Answer: %d", product)
	fmt.Println()
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
