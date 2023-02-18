package format

import (
	"time"
)

var (
	DateTime = "2006-01-02 15:04:05"
	DateOnly = "2006-01-02"
	TimeOnly = "15:04:05"
)

/*
GetTime 获取时间并格式化
timeF表示目标时间格式，t表示传入的时间
t为空时，默认使用当前时间
*/
func GetTime(timeF string, t string) (result string) {
	if t != "" {
		parseTime, _ := time.ParseInLocation(timeF, t, time.Local)
		return parseTime.String()
	} else {
		return time.Unix(time.Now().Unix(), 0).Format(timeF)
	}
}
