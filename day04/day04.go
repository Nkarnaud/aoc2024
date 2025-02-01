package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(fileName string) [][]rune {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error openning file:", err)
	}
	defer file.Close()

	var matrix [][]rune

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))

	}
	return matrix
}

func getAllColumns(matrix [][]string) [][]string {
	rows := len(matrix)
	cols := len(matrix[0])
	columns := make([][]string, cols)

	for col := 0; col < cols; col++ {
		column := make([]string, rows)
		for row := 0; row < rows; row++ {
			column[row] = matrix[row][col]
		}
		columns[col] = column
	}

	return columns
}

func getPrimaryDiagonals(matrix [][]string) [][]string {
	rows := len(matrix)
	cols := len(matrix[0])
	var diagonals [][]string

	for col := 0; col < cols; col++ {
		diagonal := []string{}
		for i, j := 0, col; i < rows && j < cols; i, j = i+1, j+1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		diagonals = append(diagonals, diagonal)
	}

	for row := 1; row < rows; row++ {
		diagonal := []string{}
		for i, j := row, 0; i < rows && j < cols; i, j = i+1, j+1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}

func getAllSecondaryDiagonals(matrix [][]string) [][]string {
	rows := len(matrix)
	cols := len(matrix[0])
	var diagonals [][]string

	for col := cols - 1; col >= 0; col-- {
		diagonal := []string{}
		for i, j := 0, col; i < rows && j >= 0; i, j = i+1, j-1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		diagonals = append(diagonals, diagonal)
	}

	for row := 1; row < rows; row++ {
		diagonal := []string{}
		for i, j := row, cols-1; i < rows && j >= 0; i, j = i+1, j-1 {
			diagonal = append(diagonal, matrix[i][j])
		}
		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}

func totalOccurrences(matrix [][]rune, word string) int {
	rows := len(matrix)
	cols := len(matrix[0])
	wordLen := len(word)
	directions := [][2]int{
		{0, 1},   // Right
		{1, 0},   // Down
		{1, 1},   // Down-Right Diagonal
		{1, -1},  // Down-Left Diagonal
		{0, -1},  // Left (Reverse Right)
		{-1, 0},  // Up (Reverse Down)
		{-1, -1}, // Up-Left Diagonal
		{-1, 1},  // Up-Right Diagonal
	}

	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if isMatch(matrix, word, r, c, dir, wordLen) {
					count++
				}
			}
		}
	}

	return count

}

func isMatch(grid [][]rune, word string, r, c int, dir [2]int, wordLen int) bool {
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < wordLen; i++ {
		nr := r + i*dir[0]
		nc := c + i*dir[1]
		if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != rune(word[i]) {
			return false
		}
	}
	return true
}

func isXMASMatch(matrix [][]rune) int {
	rows := len(matrix)
	cols := len(matrix[0])
	count := 0
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if matrix[i][j] == 'A' {
				if matrix[i-1][j-1] == 'M' && matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M' && matrix[i-1][j+1] == 'S' {
					count++
				} else if matrix[i-1][j-1] == 'S' && matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M' && matrix[i-1][j+1] == 'M' {
					count++
				} else if matrix[i-1][j-1] == 'M' && matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S' && matrix[i-1][j+1] == 'S' {
					count++

				} else if matrix[i-1][j-1] == 'S' && matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S' && matrix[i-1][j+1] == 'M' {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	matrix := readFile("input.txt")
	word := "XMAS"
	totalOccurrences := totalOccurrences(matrix, word)
	fmt.Println(totalOccurrences)
	total_count := isXMASMatch(matrix)
	fmt.Println(total_count)
}
