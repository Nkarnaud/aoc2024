package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) ([][]string, [][]string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error openning file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var page_ordering [][]string
	var page_update [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			ordering := strings.Split(line, "|")
			page_ordering = append(page_ordering, ordering)
		} else {
			update := strings.Split(line, ",")
			page_update = append(page_update, update)
		}
	}

	return page_ordering, page_update
}

func updateRule(ordering_rules [][]string) map[string][]string {
	orderMap := make(map[string][]string)
	for i := 0; i < len(ordering_rules); i++ {
		rule := ordering_rules[i]
		orderMap[rule[0]] = append(orderMap[rule[0]], rule[1])
	}
	return orderMap
}

func isValideUpdate(update []string, orderMap map[string][]string) bool {
	for i := 0; i < len(update); i++ {
		for _, dependent := range orderMap[update[i]] {
			for j := i + 1; j < len(update); j++ {
				if update[j] == dependent {
					goto Next
				}
			}
			return false
		Next:
		}
	}
	return true
}

func getMiddle(update []string) string {
	fmt.Println(update)
	return update[len(update)/2]
}

func main() {
	ordering_pages, update_pages := readFile("input.txt")
	orderMap := updateRule(ordering_pages)
	total_middle := 0
	for i := 0; i < len(update_pages); i++ {
		page := update_pages[i]
		if isValideUpdate(page, orderMap) {
			fmt.Println(getMiddle(page))

			middle, _ := strconv.Atoi(getMiddle(update_pages[i]))

			total_middle += middle
		}
	}
	fmt.Println(total_middle)
}
