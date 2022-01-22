package time

import (
	"time"
)

const ISO8601 = "2006-01-02T15:04:05.000Z"

func GetNowTime() string {
	t := time.Now()
	return t.Format(ISO8601)
}

const logFormat = "20060102150405"

func GetLogTime() string {
	t := time.Now()
	return t.Format(logFormat)
}
