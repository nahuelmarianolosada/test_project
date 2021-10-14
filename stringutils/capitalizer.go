package stringUtils

import "strings"

func Capitalize(str string) string {
	return strings.Title(str)
}

func CapitalizeAll(str string) string {
	return strings.ToUpper(str)
}
