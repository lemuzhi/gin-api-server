package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	logid, _ := c.Get("logid")
	fmt.Println(logid)
	c.JSON(200, gin.H{"message": "hello"})
}
