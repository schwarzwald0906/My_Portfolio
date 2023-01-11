package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	myerror "github.com/schwarzwald0906/My_Portfolio/src/core/myerror"
)

func ErrHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		var ERR_KEY string
		errVal, ok := c.Get(ERR_KEY)
		fmt.Println(errVal)

		val := myerror.NotFoundWrapf("error")
		fmt.Println("valの値は")
		fmt.Println(val)
		if ok {
			// エラーハンドリングをする
			switch val.(type) {
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
