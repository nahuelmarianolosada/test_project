package stringUtils

import "runtime"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func IsEmpty(s string) bool {
	return s == "" || len(s) == 0
}

func GetOSEnvironment() string {
	operatingSystem := runtime.GOOS
	if operatingSystem == "darwin" {
		return "Running in Mac OS"
	} else if operatingSystem == "linux" {
		return "Running in Linux"
	} else {
		return "Running in Windows"
	}
}
