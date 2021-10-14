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
 * Complete the 'countingValleys' function below.
 *
 *	An avid hiker keeps meticulous records of their hikes. During the last hike that took exactly steps steps, for every step it was noted if it was an uphill, U,
 *	or a downhill, D step. Hikes always start and end at sea level, and each step up or down represents a  unit change in altitude. We define the following terms:

 *  A mountain is a sequence of consecutive steps above sea level, starting with a step up from sea level and ending with a step down to sea level.
 *  A valley is a sequence of consecutive steps below sea level, starting with a step down from sea level and ending with a step up to sea level.
 *  Given the sequence of up and down steps during a hike, find and print the number of valleys walked through.

 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 *
 * Sample Input
 *  8
 *  UDDDUDUU
 * Sample Output
 *  1
 */

func countingValleys(steps int32, path string) int32 {
	pathArr := strings.Split(path, "")
	actualLevel := 0
	valleysCount := 0
	for _, v := range pathArr {
		if v == "D" {
			actualLevel--
		} else { //v == "U"
			actualLevel++
			if actualLevel == 0 {
				valleysCount++
			}
		}
	}

	return int32(valleysCount)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	steps := int32(stepsTemp)

	path := readLine(reader)

	result := countingValleys(steps, path)

	fmt.Fprintf(writer, "%d\n", result)

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
