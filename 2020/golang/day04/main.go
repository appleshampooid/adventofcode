package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"regexp"
	"strconv"
)

type Passport struct {
	BirthYear  int
	IssueYear  int
	ExpYear    int
	Height     string
	HairColor  string
	EyeColor   string
	PasspordID string
	CountryID  string
}

func NewPassport(input string) (Passport, error) {
	p := Passport{}
	fields := strings.Fields(input)
	err_prefix := "Field doesn't satisfy the format requirements"
	hair_re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	height_re := regexp.MustCompile(`^(\d+)(in|cm)$`)
	eye_re := regexp.MustCompile(`^(?:amb|blu|brn|gry|grn|hzl|oth)$`)
	pid_re := regexp.MustCompile(`^\d{9}$`)
	for _, field := range fields {
		parsed_field := strings.Split(field, ":")
		if len(parsed_field) != 2 {
			return p, fmt.Errorf("%s: wrong number of elements in %s.", err_prefix, parsed_field)
		}
		switch parsed_field[0] {
		case "byr":
			year, err := strconv.Atoi(parsed_field[1])
			if err != nil || year > 2002 || year < 1920 {
				return p, fmt.Errorf("%s: %s is not int or outside bounds", err_prefix, parsed_field[1])
			}
			p.BirthYear = year
		case "iyr":
			year, err := strconv.Atoi(parsed_field[1])
			if err != nil || year > 2020 || year < 2010 {
				return p, fmt.Errorf("%s: %s is not int", err_prefix, parsed_field[1])
			}
			p.IssueYear = year
		case "eyr":
			year, err := strconv.Atoi(parsed_field[1])
			if err != nil || year > 2030 || year < 2020 {
				return p, fmt.Errorf("%s: %s is not int", err_prefix, parsed_field[1])
			}
			p.ExpYear = year
		case "hgt":
			matches := height_re.FindStringSubmatch(parsed_field[1])
			if len(matches) != 3 {
				return p, fmt.Errorf("%s: %s not in correct height format", err_prefix, parsed_field[1])
			}
			height, _ := strconv.Atoi(matches[1])
			if matches[2] == "cm" {
				if height < 150 || height > 193 {
					return p, fmt.Errorf("%s: %d is outside of bounds", err_prefix, height)
				}
			} else {
				if height < 59 || height > 76 {
					return p, fmt.Errorf("%s: %d is outside of bounds", err_prefix, height)
				}
			}
			p.Height = parsed_field[1]
		case "hcl":
			if !hair_re.MatchString(parsed_field[1]) {
				return p, fmt.Errorf("%s: %s does not match hair color format", err_prefix, parsed_field[1])
			}
			p.HairColor = parsed_field[1]
		case "ecl":
			if !eye_re.MatchString(parsed_field[1]) {
				return p, fmt.Errorf("%s: %s does not match eye color format", err_prefix, parsed_field[1])
			}
			p.EyeColor = parsed_field[1]
		case "pid":
			// id, err := strconv.Atoi(parsed_field[1])
			// if err != nil {
			// 	return p, fmt.Errorf("%s: %s is not int", err_prefix, parsed_field[1])
			// }
			if !pid_re.MatchString(parsed_field[1]) {
				return p, fmt.Errorf("%s: %s does not match pid format", err_prefix, parsed_field[1])
			}
			p.PasspordID = parsed_field[1]
		case "cid":
			// id, err := strconv.Atoi(parsed_field[1])
			// if err != nil {
			// 	return p, fmt.Errorf("%s: %s is not int", err_prefix, parsed_field[1])
			// }
			p.CountryID = parsed_field[1]
		default:
			return p, fmt.Errorf("%s: unknown field %s", err_prefix, parsed_field[0])
		}
	}
	// fmt.Printf("%v\n", fields)
	if p.BirthYear == 0 || p.IssueYear == 0 || p.ExpYear == 0 || p.PasspordID == "" ||
		p.EyeColor == "" || p.HairColor == "" || p.Height == "" {
		return p, fmt.Errorf("Passport is missing requried fields: %v", p)
	}
	return p, nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Need an input file")
		os.Exit(1)
	}
	fileName := flag.Args()[0]
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file %s\n", fileName)
		os.Exit(1)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var records []string = make([]string, 0, 10)
	var collector string = ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			records = append(records, collector)
			collector = ""
		} else {
			collector += fmt.Sprintf(" %s", line)
		}
	}
	records = append(records, collector)
	validPassports := 0
	for _, record := range records {
		if p, err := NewPassport(record); err == nil {
			validPassports++
			fmt.Printf("Passport record: %v\n", p)
		} else {
			fmt.Printf("Error parsing record: %s\n", err.Error())
		}
	}
	fmt.Printf("There are %d valid passports.\n", validPassports)
}
