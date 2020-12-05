package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func (p passport) validPartOne() bool {
	return p.Byr != "" &&
		p.Iyr != "" &&
		p.Eyr != "" &&
		p.Hgt != "" &&
		p.Hcl != "" &&
		p.Ecl != "" &&
		p.Pid != ""
}

func (p passport) validPartTwo() bool {
	if p.validPartOne() {
		byr, byrErr := strconv.Atoi(p.Byr)
		if byrErr != nil {
			log.Fatal(byrErr)
		}
		iyr, iyrErr := strconv.Atoi(p.Iyr)
		if iyrErr != nil {
			log.Fatal(iyrErr)
		}
		eyr, eyrErr := strconv.Atoi(p.Eyr)
		if eyrErr != nil {
			log.Fatal(eyrErr)
		}
		hgtValid := false
		hgtRegexString := "(\\d+)(in|cm)"
		hgtMatched, hgtErr := regexp.MatchString(hgtRegexString, p.Hgt)
		if hgtErr != nil {
			log.Fatal(hgtErr)
		}
		if hgtMatched {
			hgtRegex := regexp.MustCompile(hgtRegexString)
			hgtRegexMatch := hgtRegex.FindStringSubmatch(p.Hgt)
			value, valueErr := strconv.Atoi(hgtRegexMatch[1])
			if valueErr != nil {
				log.Fatal(valueErr)
			}
			unit := hgtRegexMatch[2]
			if (unit == "in" && (59 <= value && value <= 76)) || (unit == "cm" && (150 <= value && value <= 193)) {
				hgtValid = true
			}
		}
		hclMatched, hclErr := regexp.MatchString("#([0-9]|[a-f]){6}", p.Hcl)
		if hclErr != nil {
			log.Fatal(hclErr)
		}
		eclMatched, eclErr := regexp.MatchString("^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$", p.Ecl)
		if eclErr != nil {
			log.Fatal(eclErr)
		}
		pidMatched, pidErr := regexp.MatchString("\\d{9}", p.Pid)
		if pidErr != nil {
			log.Fatal(pidErr)
		}
		return (1920 <= byr && byr <= 2002) &&
			(2010 <= iyr && iyr <= 2020) &&
			(2020 <= eyr && eyr <= 2030) &&
			hgtValid &&
			hclMatched &&
			eclMatched &&
			pidMatched && len(p.Pid) == 9
	}

	return false
}

func partOne(input []string) {
	validPassports := 0
	for _, pass := range input {
		p := passport{}
		fields := strings.Fields(pass)
		for _, field := range fields {
			nameValue := strings.Split(field, ":")
			f := reflect.ValueOf(&p).Elem().FieldByName(strings.Title(nameValue[0]))
			if f.IsValid() {
				f.SetString(nameValue[1])
			}
		}
		if p.validPartOne() {
			validPassports++
		}
	}

	fmt.Printf("Part 1 Answer: %d", validPassports)
	fmt.Println()
}

func partTwo(input []string) {
	validPassports := 0
	for _, pass := range input {
		p := passport{}
		fields := strings.Fields(pass)
		for _, field := range fields {
			nameValue := strings.Split(field, ":")
			f := reflect.ValueOf(&p).Elem().FieldByName(strings.Title(nameValue[0]))
			if f.IsValid() {
				f.SetString(nameValue[1])
			}
		}
		if p.validPartTwo() {
			validPassports++
		}
	}

	fmt.Printf("Part 2 Answer: %d", validPassports)
	fmt.Println()
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	passports := strings.Split(string(content), "\n\n")
	partOne(passports)
	partTwo(passports)
}
