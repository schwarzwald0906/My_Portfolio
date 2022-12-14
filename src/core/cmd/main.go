package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/controller"
)

func main() {
	// Gin のルーターを作成
	router := gin.Default()

	// ルーティング
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	controller.UserSetupRoutes(router)

	// Web サーバーを起動
	router.Run()

}

// ヘルスチェック処理
func isHealthy() bool {
	return true
}
