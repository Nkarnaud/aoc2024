package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func loadAndProcessFile(fileName string) int {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error while closing the file", err)
		}
	}()

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	buffer := bufio.NewScanner(file)
	sum := 0
	for buffer.Scan() {
		line := buffer.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 0 {
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				mul := num1 * num2
				sum += mul
			}
		}
	}
	return sum
}

func part2(fileName string) int {
	file, _ := os.Open(fileName)

	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dontRe := regexp.MustCompile(`don't\(\)`)

	buffer := bufio.NewScanner(file)
	sum := 0
	enabled := true
	for buffer.Scan() {
		line := buffer.Text()
		tokens := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`).FindAllString(line, -1)
		for _, token := range tokens {
			if doRe.MatchString(token) {
				enabled = true
			} else if dontRe.MatchString(token) {
				enabled = false
			} else if enabled && mulRe.MatchString(token) {
				match := mulRe.FindStringSubmatch(token)
				num1, _ := strconv.Atoi(match[1])
				num2, _ := strconv.Atoi(match[2])
				sum += num1 * num2
			}
		}
	}
	return sum
}

func main() {
	output := loadAndProcessFile("input.txt")
	fmt.Println(output)

	fmt.Println(part2("input.txt"))
}
