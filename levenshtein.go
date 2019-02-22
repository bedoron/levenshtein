package levenshtein

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: Variadic function ?
func getMin(v1 int, v2 int, v3 int) int {
	min := v1
	if min > v2 {
		min = v2
	}

	if min > v3 {
		return v3
	}

	return min
}

func printMat(resultMatrix [][]int, rows int, cols int) {
	for row := 0; row < rows; row++ {
		var elmsSlice []string
		for col := 0; col < cols; col++ {
			elmsSlice = append(elmsSlice, strconv.Itoa(resultMatrix[row][col]))
		}
		fmt.Println(strings.Join(elmsSlice, ","))
	}
}


func Levenshtein(s1 string, s2 string) int {
	cols := len(s1) + 1
	rows := len(s2) + 1

	resultMatrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		resultMatrix[i] = make([]int, cols)
	}

	for i := 0; i < cols; i++ {
		resultMatrix[0][i] = i
	}

	for i := 0; i < rows; i++ {
		resultMatrix[i][0] = i
	}

	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			substCost := 1
			if s1[col - 1] == s2[row - 1] {
				substCost = 0
			}

			addition := resultMatrix[row - 1][col] + 1
			deletion := resultMatrix[row][col - 1] + 1
			replace := resultMatrix[row - 1][col - 1] + substCost

			resultMatrix[row][col] = getMin(addition, deletion, replace)
		}
	}

	return resultMatrix[rows - 1][cols - 1]
}


func  main() {
	fmt.Println("Hello world")
}