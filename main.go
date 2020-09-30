package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "world",
		})
	})
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("r.run err:%s", err)
	}
}
