package main

import (
	"github.com/gin-gonic/gin"
)

func MainHttpMocks(Method string, URL string) {
	r := gin.Default()
	r.Handle(Method, URL)

}

func MockContext(c *gin.Context) {

}
