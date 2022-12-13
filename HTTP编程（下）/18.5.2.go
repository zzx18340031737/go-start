package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//静态资源加载，本例为css，以及资源图片
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFile("favicon.ico", "./static/favicon.ico")

	// Listen and serve on 0.0.0.0:80
	router.Run(":8080")
}
