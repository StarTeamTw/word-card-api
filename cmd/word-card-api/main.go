package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, GO!!")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
