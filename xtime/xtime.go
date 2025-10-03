package xtime

import "time"

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
