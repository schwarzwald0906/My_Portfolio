package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/app/userapp"
	mydatabase "github.com/schwarzwald0906/My_Portfolio/src/core/infra/database"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/repoimpl"
)

// ルーティングを別の関数やメソッドに分割
func UserSetupRoutes(router *gin.Engine) {
	router.GET("/users", func(c *gin.Context) {
		// 正常系
		if isHealthy() {
			c.JSON(http.StatusOK, gin.H{"data": "user list"})
			return
		}
		// 異常系
		c.String(http.StatusServiceUnavailable, "unavailable")

	})

	router.POST("/users", func(c *gin.Context) {
		// 正常系
		if isHealthy() {
			c.String(http.StatusOK, "ok")
			return
		}
		// 異常系
		c.String(http.StatusServiceUnavailable, "unavailable")

		//データベース接続
		repo := mydatabase.DbInit()
		userRepo := repoimpl.NewUserRepository(repo)

		// コンストラクタ作成
		createUserApp := userapp.NewCreateUserApp(userRepo)

		// フォームからデータを取得
		req := &userapp.CreateUserRequest{
			Email:    c.PostForm("mail"),
			Password: c.PostForm("password"),
		}
		// 一旦ハードコーディング
		req.Email = "email@gmail.com"
		req.Password = "password12345!"

		if err := createUserApp.Exec(c, req); err != nil {
			c.AbortWithStatus(500)
			return
		}
		// レスポンスを返す
		c.JSON(201, nil)

		//データベースの接続を切る
		defer repo.Close()
	})
}
