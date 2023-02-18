package middleware

import (
	"gin-project-template/global"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

// AccessLogger 日志中间件
func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logid := strings.Join([]string{"520", strconv.FormatInt(time.Now().UnixNano(), 10)}, "")
		c.Set("logid", logid)

		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		end := time.Now().Sub(start) //计算请求耗时

		body := make([]byte, c.Request.ContentLength)
		_, err := c.Request.Body.Read(body)
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		var param string
		if c.Request.ContentLength > 0 {
			param = strings.ReplaceAll(string(body), "\r", "")
			param = strings.ReplaceAll(param, "\n", "")
			param = strings.ReplaceAll(param, " ", "")
			//if c.Request.Header.Get("content-type") == binding.MIMEMultipartPOSTForm {
			//	fmt.Println("content-type")
			//	param = strings.Join(regexp.MustCompile(`form-data;(.*?)----------`).FindAllString(param, -1), "")
			//}
		}
		if raw != "" {
			path = path + "?" + raw
		}
		v := map[string]interface{}{
			"method":  c.Request.Method,
			"status":  c.Writer.Status(),
			"latency": end,
			"ip":      c.ClientIP(),
			//"time:", start.Format(time.RFC850),
			"path":         path,
			"content-type": c.Request.Header.Get("content-type"),
			"param":        param,
			"logid":        logid,
			"errors":       c.Errors.String(),
			"user-agent":   c.Request.UserAgent(),
		}
		//data := common.MapToJson(v)
		//fmt.Println("数据", data)
		global.Log.InfoJson(v)
		//global.Log.Infof(`{"method":%s,"status":%d,"latency":%v,"ip":%s,"path":%s,"content-type:%v,"param":%v,"logid":%s,"errors":%v,"user-agent":%v}`,
		//	c.Request.Method,
		//	c.Writer.Status(),
		//	end,
		//	c.ClientIP(),
		//	path,
		//	c.Request.Header.Get("content-type"),
		//	param,
		//	logid,
		//	c.Errors.String(),
		//	c.Request.UserAgent(),
		//)
		//err = log.Output(1, fmt.Sprintf(" |%v %3d %8v| %13v | %v | %v | %v\n",
		//	c.Request.Method,
		//	c.Writer.Status(),
		//	end,
		//	c.ClientIP(),
		//	//"time:", start.Format(time.RFC850),
		//	path,
		//	c.Request.Header.Get("content-type"),
		//
		//	//"errors: ", c.Errors.String(),
		//	param,
		//	//"user-agent:", c.Request.UserAgent(),
		//))
		//if err != nil {
		//	log.Println(err)
		//}
	}
}
