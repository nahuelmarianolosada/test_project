package hackerrank

import (
	"fmt"
	"strconv"
	"strings"
)

/**
* Implement a function thah co7unds the number of set bits in the binary representation of a 32-bit integer
 */
func main() {
	fmt.Println("Cantidad de bits: " + fmt.Sprintf("%d", countBits(uint32(126))))
}

/*
 * Complete the 'countBits' function below.
 *
 * The function is expected to return an int32.
 * The function accepts unit32 num as parameter.
 */

func countBits(num uint32) int32 {
	totalCountBits := strconv.FormatInt(int64(num), 2)

	return int32(strings.Count(totalCountBits, "1"))
}
