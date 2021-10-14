package conversions

import (
	"strconv"
)

func BoolToString(x bool) string {
	return strconv.FormatBool(x)
}
