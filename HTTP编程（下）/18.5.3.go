package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := c.Request.Method + "=>" + name + " is " + action

		c.String(http.StatusOK, message)
	})

	router.POST("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := c.Request.Method + "=>" + name + " is " + action

		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
