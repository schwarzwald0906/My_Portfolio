package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schwarzwald0906/My_Portfolio/src/core/app/userapp"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/userdm"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/repoimpl"
)

func main() {
	// Gin のルーターを作成
	r := gin.Default()

	// ユーザー作成ボタンを押下したときの処理
	r.POST("/users", func(c *gin.Context) {

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

	// Web サーバーを起動
	r.Run()
}

// ヘルスチェック処理
func isHealthy() bool {
	return true
}
