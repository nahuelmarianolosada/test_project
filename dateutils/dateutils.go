package dateUtils

import "time"

const (
	SHORT_FORMAT = "02-01-2006"
	LONG_FORMAT  = SHORT_FORMAT + " 15:04:00"
)

func ToString(date time.Time) string {
	return date.String()
}

func FormatShort(date time.Time) string {
	return date.Format(SHORT_FORMAT)
}

func FormatLong(date time.Time) string {
	return date.Format(LONG_FORMAT)
}

func AddMonths(date time.Time, months int) time.Time {
	return date.AddDate(0, months, 0)
}
