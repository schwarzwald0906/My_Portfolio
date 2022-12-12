package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/app/userapp"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/userdm"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/repoimpl"
)

// ルーティングを別の関数やメソッドに分割
func SetupRoutes(router *gin.Engine) {
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

		var userRepo userdm.UserRepository = repoimpl.NewUserRepository(&sql.DB{})
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
	})
}

// ヘルスチェック処理
func isHealthy() bool {
	return true
}

// ユーザー用のコントローラーを作成する
// ctrl := UserController{}

// v1 := engine.Group("/v1")
// {
// 	// コントローラを通して、出力したい場合。
// 	v1.GET("/login", ctrl.login)
// 	v1.GET("/signup", ctrl.signup)
// 	v1.POST("/signup", ctrl.signup)
// }
