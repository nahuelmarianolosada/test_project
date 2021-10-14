package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * There is a large pile of socks that must be paired by color.
 * Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n => the number of socks in the pile
 *  2. INTEGER_ARRAY ar => the colors of each sock
 */

func SockMerchant(n int32, ar []int32) int32 {
	result := map[int32]int32{}
	sockWithMatchig := int32(0)
	for _, v := range ar {
		if _, ok := result[v]; !ok {
			result[v] = 1
		} else {
			result[v]++
		}
	}
	r, _ := json.Marshal(&result)
	fmt.Printf("%s", string(r))
	for _, v := range result {
		sockWithMatchig += v / 2
	}
	return int32(sockWithMatchig)
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	// nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// checkError(err)
	// n := int32(nTemp)

	// arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	// var ar []int32

	// for i := 0; i < int(n); i++ {
	// 	arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
	// 	checkError(err)
	// 	arItem := int32(arItemTemp)
	// 	ar = append(ar, arItem)
	// }

	// result := sockMerchant(n, ar)
	fmt.Println(SockMerchant(6, []int32{1, 2, 3, 4, 5, 6}))
	// fmt.Fprintf(writer, "%d\n", result)

	// writer.Flush()
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
