package remaindersorting

import (
	"fmt"
	"sort"
)

/*
*
* Implement a function that receives an array of strings and sorts them based on the following heuristics:
* - The primary sort is by increasing remainder of the strings' lenghts, modulo 3
* - If multiple strings have the same remainder, they should be sorted in alphabetical order.
*
* Example:
* strArr = ["Colorado", "Utah", "Wisconsin", "Oregon"]
* Their lenghts are [8, 4, 9, 6]. The remainders, modulo 3, are [2, 1, 0, 0]. Oregon and Wisconsin have the
* same remainder and are further sorted sorted alphabetically.
* Here is the sorted array showing (string, length, length modulo):
* [("Oregon", 6, 0), ("Wisconsin", 9, 0), ("Utah", 4, 1), ("Colorado",8, 2)]
* Return the sorted array ["Oregon", "Wisconsin", "Utah", "Colorado"]
*
 */
func main() {
	fmt.Println(fmt.Sprintf("%v", RemainderSorting([]string{"Colorado", "Utah", "Wisconsin", "Oregon"})))
}

func RemainderSorting(strArr []string) []string {

}

func fillDataStrings(strArr []string) (returnArr []StringsWithData) {
	for _, v := range strArr {
		newStringWithData := StringsWithData{
			Name:      v,
			Length:    len(v),
			LengthMod: len(v) % 3,
		}
		returnArr = append(returnArr, newStringWithData)
	}
	return returnArr
}

func sortDataByLength(dataArr []StringsWithData) []StringsWithData {
	sort.SliceStable(dataArr, func(i, j int) bool {
		return dataArr[i].Length < dataArr[j].Length
	})
	return dataArr
}

type StringsWithData struct {
	Name      string
	Length    int
	LengthMod int
}
