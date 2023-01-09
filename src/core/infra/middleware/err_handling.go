package middleware

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/error"
)

func ErrHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		var ERR_KEY string
		errVal, ok := c.Get(ERR_KEY)
		fmt.Println(errVal)

		val := error.BadRequestWrapf("error")
		fmt.Println("valの値は")
		fmt.Println(val)
		if ok {
			// エラーハンドリングをする
			switch val.(type) {
			case error.BadRequestErr:
				c.AbortWithStatus(400)
			case error.NotFoundErr:
				c.AbortWithStatus(404)
			case error.InternalServerErr:
				c.AbortWithStatus(500)
			}
		}
		log.Println("after logic")
	}
}
