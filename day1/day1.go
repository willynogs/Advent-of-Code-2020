package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func partOne(input []string) {
	for x := 0; x < (len(input) - 2); x++ {
		for y := (x + 1); y < len(input); y++ {
			xI, xErr := strconv.Atoi(input[x])
			if xErr != nil {
				log.Fatal(xErr)
			}
			yI, yErr := strconv.Atoi(input[y])
			if yErr != nil {
				log.Fatal(yErr)
			}
			if xI+yI == 2020 {
				fmt.Printf("Part 1 Answer: %d", xI*yI)
				fmt.Println()
			}
		}
	}
}

func partTwo(input []string) {
	for x := 0; x < (len(input) - 2); x++ {
		for y := (x + 1); y < (len(input) - 1); y++ {
			for z := (y + 1); z < len(input); z++ {
				xI, xErr := strconv.Atoi(input[x])
				if xErr != nil {
					log.Fatal(xErr)
				}
				yI, yErr := strconv.Atoi(input[y])
				if yErr != nil {
					log.Fatal(yErr)
				}
				zI, zErr := strconv.Atoi(input[z])
				if zErr != nil {
					log.Fatal(zErr)
				}
				if xI+yI+zI == 2020 {
					fmt.Printf("Part 2 Answer: %d", xI*yI*zI)
					fmt.Println()
				}
			}
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
