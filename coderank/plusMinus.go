package coderank



import (
"bufio"
"os"
"strconv"
"strings"
	"fmt"
)

func GetPositiveNegativeAndZero(arr []int32) (float64, float64, float64) {
	var posittives, negatives, zeros float64
	for _, v := range arr {
		switch {
			case v == 0 : zeros++
			case v > 0 : posittives++
			case v < 0 : negatives++
		}
	}
	return posittives, negatives, zeros
}

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	pos, neg, zer := GetPositiveNegativeAndZero(arr)

	fmt.Printf("%f\n", float64(pos)/float64(len(arr)))
	fmt.Printf("%f\n", float64(neg)/float64(len(arr)))
	fmt.Printf("%f\n", float64(zer)/float64(len(arr)))
}

func mainPlusMinus() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

