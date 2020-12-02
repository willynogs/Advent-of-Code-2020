package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type password struct {
	text         string
	requiredChar string
	requiredMin  int
	requiredMax  int
}

func parsePassword(input string) *password {
	arr := strings.Split(input, " ")
	minMax := strings.Split(arr[0], "-")
	min, minErr := strconv.Atoi(minMax[0])
	if minErr != nil {
		log.Fatal(minErr)
	}
	max, maxErr := strconv.Atoi(minMax[1])
	if maxErr != nil {
		log.Fatal(maxErr)
	}
	return &password{requiredMin: min, requiredMax: max, requiredChar: strings.ReplaceAll(arr[1], ":", ""), text: arr[2]}
}

func passwordValidPart1(input *password) bool {
	appearances := strings.Count(input.text, input.requiredChar)
	valid := input.requiredMin <= appearances && appearances <= input.requiredMax
	return valid
}

func passwordValidPart2(input *password) bool {
	min := input.requiredMin - 1
	max := input.requiredMax - 1
	if min < len(input.text) && max < len(input.text) {
		containsMin := string(input.text[min]) == input.requiredChar
		containsMax := string(input.text[max]) == input.requiredChar
		return (containsMin && !containsMax) || (!containsMin && containsMax)
	}
	return false
}

func partOne(input []string) {
	validPasswords := 0
	for i := 0; i < len(input); i++ {
		if passwordValidPart1(parsePassword(input[i])) {
			validPasswords++
		}
	}
	fmt.Printf("Part 1 Answer: %d", validPasswords)
	fmt.Println()
}

func partTwo(input []string) {
	validPasswords := 0
	for i := 0; i < len(input); i++ {
		if passwordValidPart2(parsePassword(input[i])) {
			validPasswords++
		}
	}
	fmt.Printf("Part 2 Answer: %d", validPasswords)
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
