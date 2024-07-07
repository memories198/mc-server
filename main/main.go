package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"mc-server/consts"
)

func main() {
	r := gin.Default()

	// 定义文件下载的路由
	r.GET("/download/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		filePath := consts.ResourcePacks + filename

		// 设置响应头
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")

		// 发送文件
		c.File(filePath)
	})

	// 运行服务器
	err := r.Run(":32745")
	if err != nil {
		log.Println(err)
		return
	}
}
