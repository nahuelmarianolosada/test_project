package stringUtils

import "strings"

func Replace(s string, old string, new string, matches int) string {
	return strings.Replace(s, old, new, matches)
}
