package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	entries := strings.Split(string(content), "\n")

	for x := 0; x < (len(entries) - 1); x++ {
		for y := (x + 1); y < len(entries[x:]); y++ {
			xI, xErr := strconv.Atoi(entries[x])
			if xErr != nil {
				log.Fatal(xErr)
			}
			yI, yErr := strconv.Atoi(entries[y])
			if yErr != nil {
				log.Fatal(yErr)
			}
			if xI+yI == 2020 {
				fmt.Printf("Answer: %d", xI*yI)
			}
		}
	}
}
