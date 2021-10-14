package coderank

import (
"bufio"
"os"
"strconv"
"strings"
	"sort"
	"fmt"
)

// Complete the miniMaxSum function below.
// Given five positive integers,
// find the minimum and maximum values that can be calculated by summing exactly four of the five integers.
// Then print the respective minimum and maximum values as a single line of two space-separated long integers.
//
//Example
//arr := []int32{1,3,5,7,9}
//The minimum sum is {1,3,5,7}=16 and the maximum sum is {3,5,7,9}=24.
//
// The function prints
//16 24
func MiniMaxSum(arr []int32) {
	arrInt := parseArrInt32ToInt(arr)
	sort.Ints(arrInt)
	var sumMini, sumMax int
	for i, v := range arrInt {
		if i != len(arr) - 1 { sumMini += v }
		if i != 0 { sumMax += v }
	}
	fmt.Printf("%d %d", sumMini, sumMax)
}

func parseArrInt32ToInt(arr []int32) (returnArr []int) {
	for _,v := range arr {
		returnArr = append(returnArr, int(v))
	}
	return returnArr
}

func mainMiniMax() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	MiniMaxSum(arr)
}

