package main

import (
	"github.com/gin-gonic/gin"
)

type Operators struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", hello)
	r.POST("/divide", divide)
	r.Run(":8081")
}

func hello(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{
		"massage": "Hello, " + name})
}

func divide(c *gin.Context) {
	var req Operators
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	result := req.A / req.B
	c.JSON(200, gin.H{
		"result": result,
	})
}
