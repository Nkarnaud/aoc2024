package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, []int) {
	// Open input file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file", err)
	}

	// close file on exit and check return error
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error while closing the file", err)
		}
	}()

	// buffer reader
	buffer := bufio.NewScanner(file)

	var firstList []int
	var secondList []int

	for buffer.Scan() {
		line := buffer.Text()
		numbers := strings.Split(line, "   ")
		num_1, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println(err)
		}
		num_2, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println(err)
		}
		firstList = append(firstList, num_1)
		secondList = append(secondList, num_2)
	}
	return firstList, secondList
}

func quickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	pivot := list[len(list)-1]
	leftList := []int{}
	rightList := []int{}

	for _, num := range list[:len(list)-1] {
		if num <= pivot {
			leftList = append(leftList, num)
		} else {
			rightList = append(rightList, num)
		}
	}

	leftList = quickSort(leftList)
	rightList = quickSort(rightList)

	return append(append(leftList, pivot), rightList...)
}

func sumList(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calculateSimilarity(number int, arr []int) int {
	if len(arr) < 1 {
		return 0
	}

	occurrence := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == number {
			occurrence += 1
		}
	}

	similarity := number * occurrence
	return similarity
}

func main() {
	file_name := "inputs.txt"
	firstList, secondList := readFile(file_name)
	sortedFirstList := quickSort(firstList)
	sortedSecondList := quickSort(secondList)

	var extraList []int
	if len(sortedFirstList) > len(sortedSecondList) {
		extraList = sortedFirstList[len(sortedSecondList):]
		sortedFirstList = sortedFirstList[:len(sortedSecondList)]
	} else {
		extraList = sortedSecondList[len(sortedFirstList):]
		sortedSecondList = sortedSecondList[:len(sortedSecondList)]
	}

	total_distance := 0
	for i := 0; i < len(sortedFirstList); i++ {
		diff := sortedFirstList[i] - sortedSecondList[i]
		if diff < 0 {
			diff = -diff
		}
		total_distance += diff
	}

	extraSum := sumList(extraList)
	result := total_distance + extraSum
	fmt.Println(result)

	// part two

	similarityScore := 0
	for _, num := range sortedFirstList {
		similarity := calculateSimilarity(num, sortedSecondList)
		similarityScore += similarity
	}

	fmt.Println(similarityScore)

}
