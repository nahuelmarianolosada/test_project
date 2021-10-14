package coderank

import (
"bufio"
"fmt"
"os"
	"time"
)

/*
 * Complete the timeConversion function below.
 * Given a time in -hour AM/PM format, convert it to military (24-hour) time.
 *
 * Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
 * - 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.
 * Example
 *
 * s := "12:01:00PM" Return '12:01:00'.
 * s := "12:01:00AM" Return '00:01:00'.
 *
 */
func TimeConversion(s string) string {
	layout := "03:04:05PM"
	layout2 := "15:04:05"
	t, err := time.Parse(layout, s)

	checkError(err)

	return t.Format(layout2)
}

func mainTimeConversion() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024 * 1024)

	s := readLine(reader)

	result := TimeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

