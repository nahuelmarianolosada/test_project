package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getTotalX' function below.
 *
 * There will be two arrays of integers. Determine all integers that satisfy the following two conditions:
 *
 *  1 - The elements of the first array are all factors of the integer being considered
 *  2 - The integer being considered is a factor of all elements of the second array
 * These numbers are referred to as being between the two arrays. Determine how many such numbers exist.
 *
 * The function is expected to return an INTEGER. It should return the number of integers that are betwen the sets.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 *
 * Example
 * a=[2,6]
 * b=[24,36]
 *
 * There are two numbers between the arrays: 6 and 12
 * 6%2=0 ; 6%6=0 ; 24%6=0  and 36%6=0 for the first value.
 * 12%2=0 ; 12%6=0 and 24%12=0 ; 36%12=0 for the second value. Return 2
 *
 * Explanation
 * 2 and 4 divide evenly into 4, 8, 12 and 16.
 * 4, 8 and 16 divide evenly into 16, 32, 96.
 *
 * 4, 8 and 16 are the only three numbers for which each element of a is a factor and each is a factor of all elements of b.
 *
 */

func areAllFactorsFor(a []int32, valueConsidered int32) bool {
	for _, v := range a {
		if valueConsidered%v != 0 {
			return false
		}
	}
	return true
}

func valueIsFactorOfAllElements(valueConsidered int32, b []int32) bool {
	for _, v := range b {
		if v%valueConsidered != 0 {
			return false
		}
	}
	return true
}

func getTotalX(a []int32, b []int32) int32 {
	total := 0

	for i := int32(1); i <= 100; i++ {
		if areAllFactorsFor(a, i) && valueIsFactorOfAllElements(i, b) {
			total++
		}
	}

	return int32(total)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
