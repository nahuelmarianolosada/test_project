package stringUtils

import "strings"

//Our spaces from the beginning and end, including the tab, are gone
func Trimmer(s string) string {
	return strings.TrimSpace(s)
}
