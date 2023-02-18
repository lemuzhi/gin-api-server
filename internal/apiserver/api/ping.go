package api

import (
	"fmt"
	"gin-project-template/global"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	logid, _ := c.Get("logid")
	global.Log.Info("ping一下, logid: ", logid)
	fmt.Println(logid)
	c.JSON(200, gin.H{"message": "hello"})
}
