package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/hello", hello)

}

func hello(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{
		"massage": "Hello, " + name})
}
