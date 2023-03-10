package internal

import (
	"gin-api-server/conf"
	"gin-api-server/initialize"
	"gin-api-server/internal/api"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"os"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Gin!</h1>
</body>
</html>
`))

func InitRouter() *gin.Engine {
	// 记录到文件。
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//var r *gin.Engine
	if conf.Server.Mode == "debug" {
		gin.DefaultWriter = io.MultiWriter(os.Stdout)
	}
	r := gin.Default()

	r.Use(initialize.GinLogger())
	r.Use(initialize.GinRecovery(true))

	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(200, "https", gin.H{
			"status": "success",
		})
	})

	router := r.Group("/api/v1")
	{
		router.POST("/ping", api.Ping)
	}

	//authRouter := r.Group("/api/v1")
	//{
	//	//TODO
	//}
	return r
}
