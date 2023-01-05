package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/controller"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/middleware"
)

func main() {
	// Gin のルーターを作成
	r := gin.Default()

	//デフォルトルーティング
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	//管理者ログイン時のルーティング処理をグループ化
	g := r.Group("/")
	g.Use(middleware.ErrHandling())
	{
		controller.UserSetupRoutes(g)
	}

	// Web サーバーを起動
	r.Run()

}
