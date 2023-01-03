package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		log.Println("after logic")
	}
}
