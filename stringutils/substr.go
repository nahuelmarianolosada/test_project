package stringUtils

func SubStr(s string, begin int, end int) string {
	return s[begin:end]
}

func SubStrWithoutEnd(s string, begin int) string {
	return s[begin:]
}
