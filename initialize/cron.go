package initialize

import (
	"fmt"
	"time"
)

func InitCron() {
	//ticker := time.NewTicker(11400 * time.Hour) // 每10天执行一次
	go func() {
		ticker := time.NewTicker(20 * time.Second) // 每1小时执行一次
		for range ticker.C {
			fmt.Println("执行")
			InitLogger()
		}
	}()
}

//func a() {
//	//每日零点定时日志回滚分割实现时间上的分割
//	if "daily" == "daily" {
//		go  func() {
//			for {
//				nowTime := time.Now()
//				nowTimeStr := nowTime.Format("2006-01-02")
//				//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
//				t2, _ := time.ParseInLocation("2006-01-02", nowTimeStr, time.Local)
//				// 第二天零点时间戳
//				next := t2.AddDate(0, 0, 1)
//				after := next.UnixNano() - nowTime.UnixNano() - 1
//				<-time.After(time.Duration(after) * time.Nanosecond)
//				loggerWrite.Rotate()
//			}
//		}()
//	}
//}
