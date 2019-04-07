package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kanziw/learn-go/01_gin/constant"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Gin is running w/ %s", constant.Version)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
