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

	for x := 0; x < (len(entries) - 2); x++ {
		for y := (x + 1); y < (len(entries) - 1); y++ {
			for z := (y + 1); z < len(entries); z++ {
				xI, xErr := strconv.Atoi(entries[x])
				if xErr != nil {
					log.Fatal(xErr)
				}
				yI, yErr := strconv.Atoi(entries[y])
				if yErr != nil {
					log.Fatal(yErr)
				}
				zI, zErr := strconv.Atoi(entries[z])
				if zErr != nil {
					log.Fatal(zErr)
				}
				if xI+yI+zI == 2020 {
					fmt.Printf("Answer: %d", xI*yI*zI)
					fmt.Println()
				}
			}
		}
	}
}
