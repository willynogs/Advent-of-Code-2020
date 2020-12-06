package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"

	"github.com/mpvl/unique"
)

func partOne(groups []string) {
	total := 0
	for _, group := range groups {
		var groupYes []string
		members := strings.Split(group, "\n")
		for _, member := range members {
			groupYes = append(groupYes, strings.Split(member, "")...)
		}
		sort.Strings(groupYes)
		unique.Strings(&groupYes)
		total += len(groupYes)
	}
	fmt.Printf("Part 1 Answer: %d", total)
	fmt.Println()
}

func partTwo(groups []string) {
	total := 0
	for _, group := range groups {
		members := strings.Split(group, "\n")
		for _, member := range members {
			strings.Split(member, "")
		}
		total += len(groupYes)
	}
	fmt.Printf("Part 2 Answer: %d", total)
	fmt.Println()
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	groups := strings.Split(string(content), "\n\n")
	partOne(groups)
	partTwo(groups)
}
