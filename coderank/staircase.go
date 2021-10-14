package coderank

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

// Complete the staircase function below.
func Staircase(n int32) {
	//ESTO LO HACE AL REVES
	//for i:=0; i<int(n); i++ {
	//	for j:=0; j<int(n); j++ {
	//		fmt.Print("# ")
	//		if j>=i {break}
	//	}
	//	fmt.Println("")
	//}
	for i := 1; int32(i) <= n; i++ {
		pt := n - int32(i)
		for c := 1; int32(c) <= pt ; c++ {
			fmt.Print(" ")
		}
		for c := 1; c <= i; c++ {
			fmt.Print("#")
		}
		if int32(i) < n {
			fmt.Println("")
		}
	}


}

func mainStaircase() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	Staircase(n)
}
