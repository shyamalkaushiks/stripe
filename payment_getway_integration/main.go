package main

import (
	services "hello/Services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	services.ConnectDatabase()
	routes := &services.HandlerService{}
	routes.Bootstrap(r)

	r.Run(":8080")
}
