package middleware

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/schwarzwald0906/My_Portfolio/src/core/config"
	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

func ErrHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		errVal, ok := c.Get(config.ERR_KEY)
		if ok {
			// エラーハンドリングをする
			switch errVal.(type) {
			case *myerror.BadRequestErr:
				c.AbortWithStatus(400)
			case *myerror.NotFoundErr:
				c.AbortWithStatus(404)
			case *myerror.InternalServerErr:
				c.AbortWithStatus(500)
			}
		}
		log.Println("after logic")
	}
}
