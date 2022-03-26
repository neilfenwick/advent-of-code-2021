package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open input for reading: %v", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	validCount, invalidCount := day4(f)

	fmt.Printf("Valid passports: %d\n", validCount)
	fmt.Printf("Invalid passports: %d\n", invalidCount)
}

func day4(reader io.Reader) (validCount int, invalidCount int) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(emptyLineSplitFunc)
	for scanner.Scan() {
		s := strings.Replace(scanner.Text(), "\n", " ", -1)
		passport, err := extractPassportFields(s)

		if err == nil && passport.isValid() {
			validCount++
		} else {
			invalidCount++
		}
	}

	return validCount, invalidCount
}

func extractPassportFields(s string) (passport, error) {
	byr, err := extractFieldInt(s, "byr:")
	if err != nil {
		return passport{}, err
	}

	iyr, err := extractFieldInt(s, "iyr:")
	if err != nil {
		return passport{}, err
	}

	eyr, err := extractFieldInt(s, "eyr:")
	if err != nil {
		return passport{}, err
	}

	hgt, err := extractField(s, "hgt:")
	if err != nil {
		return passport{}, err
	}

	hcl, err := extractField(s, "hcl:")
	if err != nil {
		return passport{}, err
	}

	ecl, err := extractField(s, "ecl:")
	if err != nil {
		return passport{}, err
	}

	pid, err := extractField(s, "pid:")
	if err != nil {
		return passport{}, err
	}

	cid, _ := extractField(s, "cid:")

	return passport{byr: byr, iyr: iyr, eyr: eyr, hgt: hgt, hcl: hcl, ecl: ecl, pid: pid, cid: cid}, nil
}

func extractField(s string, pattern string) (string, error) {
	idx := strings.Index(s, pattern)
	if idx == -1 {
		return "", fmt.Errorf("field '%v' not found", pattern)
	}

	subStr := s[idx+4:]
	spaceIdx := strings.Index(subStr, " ")
	if spaceIdx > -1 {
		return subStr[:spaceIdx], nil
	}

	return subStr, nil
}

func (p *passport) isValid() bool {
	if strings.Index(p.hgt, "in") < 1 && strings.Index(p.hgt, "cm") < 1 {
		return false
	}

	if strings.Index(p.hgt, "in") > -1 {
		hgtInches, err := strconv.ParseInt(strings.ReplaceAll(p.hgt, "in", ""), 10, 0)
		if err != nil || hgtInches < 59 || hgtInches > 76 {
			return false
		}
	}

	if strings.Index(p.hgt, "cm") > -1 {
		hgtCm, err := strconv.ParseInt(strings.ReplaceAll(p.hgt, "cm", ""), 10, 0)
		if err != nil || hgtCm < 150 || hgtCm > 193 {
			return false
		}
	}

	if p.byr < 1920 || p.byr > 2002 {
		return false
	}

	if p.iyr < 2010 || p.iyr > 2020 {
		return false
	}

	if p.eyr < 2020 || p.eyr > 2030 {
		return false
	}

	match, errHcl := regexp.MatchString("^#([0-9]|[a-f]){6}$", p.hcl)
	if !match || errHcl != nil {
		return false
	}

	match, errEcl := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", p.ecl)
	if !match || errEcl != nil {
		return false
	}

	match, errPid := regexp.MatchString(`^\d{9}$`, p.pid)
	if !match || errPid != nil {
		return false
	}

	return true
}

func extractFieldInt(s string, pattern string) (int, error) {
	str, err := extractField(s, pattern)
	if err != nil {
		return -1, err
	}

	result, err := strconv.Atoi(str)
	if err != nil {
		return -1, err
	}
	return result, nil
}

func emptyLineSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if atEOF {
		return len(data), data, nil
	}

	// Find two newlines in a row and slice out the data
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 2, data[0:i], nil
	}

	return
}
