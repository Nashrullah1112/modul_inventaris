package Constant

import "time"

const (
	DatabaseError     = "database_error"
	InternalHttpError = "internal_http_error"
)

func ParseTime(timeStr string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}
	}
	return parsedTime
}
