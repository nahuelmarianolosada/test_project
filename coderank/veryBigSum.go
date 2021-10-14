package coderank


import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

// Complete the aVeryBigSum function below.
func AVeryBigSum(ar []int64) int64 {
	var total int64
	for _, v := range ar {
		total += v
	}
	return total
}


func mainVeryBigSum() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int64

	for i := 0; i < int(arCount); i++ {
		arItem, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		ar = append(ar, arItem)
	}

	result := AVeryBigSum(ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

