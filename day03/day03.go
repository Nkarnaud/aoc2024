package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func loadFile(fileName string) []string {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error while closing the file", err)
		}
	}()

	expression := `^mul$begin:math:text$\\d{1,3},\\d{1,3}$end:math:text$$`
	re := regexp.MustCompile(expression)

	buffer := bufio.NewScanner(file)
	var array []string
	for buffer.Scan() {
		line := buffer.Text()
		matches := re.FindAllString(line, -1)
		for _, match := range matches {
			array = append(array, match)
		}

	}
	return array
}
