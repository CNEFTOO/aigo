package xtime

import (
	"strconv"
	"strings"
	"time"
)

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func GetTimestampMilli() int64 {
	return time.Now().UnixMilli()
}

func GetTimestampNano() int64 {
	return time.Now().UnixNano()
}

func GetTimestampMicro() int64 {
	return time.Now().UnixMicro()
}

func GetTimestampUTC() int64 {
	return time.Now().UTC().Unix()
}

func FormatDate(date string) string {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}
	return t.Format("2006-01-02")
}

func ConvTime(date string) string {
	return strings.Replace(date, "-", "", -1)
}

func ConvTime2Float64(date string) float64 {
	val, err := strconv.ParseFloat(date, 64)
	if err != nil {
		return 0.0
	}
	return val
}
