package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func evaluateSafty(arr []string) bool {
	if len(arr) < 2 {
		return true
	}
	first_num, err := strconv.Atoi(arr[0])
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	second_num, err := strconv.Atoi(arr[1])
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	isIncreasing := second_num > first_num
	for i := 1; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		num2, err := strconv.Atoi(arr[i-1])
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		diff := int(math.Abs(float64(num - num2)))
		if diff < 1 || diff > 3 {
			return false
		}
		if (isIncreasing && num < num2) || (!isIncreasing && num > num2) {
			return false
		}
	}
	return true
}

func canBeMadeSafe(report []string) bool {
	for i := 0; i < len(report); i++ {
		modified := append(report[:i], report[i+1:]...)
		if evaluateSafty(modified) {
			return true
		}
	}
	return false
}

func processIput(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error openning file", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error while closing file")
		}
	}()

	scanner := bufio.NewScanner(file)

	safeReport := 0

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		isSafed := evaluateSafty(arr)
		if isSafed {
			safeReport += 1
		} else {
			if canBeMadeSafe(arr) {
				safeReport += 1
			}
		}

	}
	return safeReport
}

func main() {
	data := processIput("input.txt")
	fmt.Println(data)
}
