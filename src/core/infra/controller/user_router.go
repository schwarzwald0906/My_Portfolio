package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/app/userapp"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/userdm"
	mydatabase "github.com/schwarzwald0906/My_Portfolio/src/core/infra/database"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/repoimpl"
)

// ルーティングを別の関数やメソッドに分割
func UserSetupRoutes(router *gin.Engine) {
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "user list"})
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
		var userRepo userdm.UserRepository = repoimpl.NewUserRepository(repo)

		// コンストラクタ作成
		createUserApp := userapp.NewCreateUserApp(userRepo)

		var req *userapp.CreateUserRequest
		// フォームからデータを取得
		// 一旦ハードコーディング
		req.Email = c.PostForm("email")
		req.Password = c.PostForm("password")

		err := createUserApp.Exec(c, req)
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		// レスポンスを返す
		c.JSON(201, nil)

		//データベースの接続を切る
		defer repo.Close()
	})
}

// ヘルスチェック処理
func isHealthy() bool {
	return true
}
