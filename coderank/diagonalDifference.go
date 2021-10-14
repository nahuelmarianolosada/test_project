package coderank

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func sumLeftToRigDiagonal(arr [][]int32) int32 {
	var sum int32
	for col := 0; col < len(arr); col++ {
		for row := 0; row < len(arr); row++ {
			if col == row {
				sum += arr[col][row]
			}
		}
	}
	return sum
}

func sumRigToLeftDiagonal(arr [][]int32) int32 {
	var sum int32
	for col := 0; col < len(arr); col++ {
		for row := 0; row < len(arr); row++ {
			if (row + col) == (len(arr) - 1) {
				sum += arr[row][col]
			}
		}
	}
	return sum
}

// Abs returns the absolute value of x.
func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func DiagonalDifference(arr [][]int32) int32 {
	return abs(sumLeftToRigDiagonal(arr) - sumRigToLeftDiagonal(arr))
}

func mainDiagonalDifference() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := DiagonalDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}
