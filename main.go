package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ymdd1/mytweet/crypto"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //直接的な記述が無いが、インポートしたいものに対しては"_"を頭につける決まり
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

// Blog モデルの宣言
type Blog struct {
	gorm.Model
	Title      string `form:"title" binding:"required"`
	Content    string `form:"content" binding:"required"`
	PictureUrl string `form:"pictureUrl" binding:"required"`
}

// User モデルの宣言
type User struct {
	gorm.Model
	Mail     string `form:"mail" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "Velden5425!"
	DBNAME := "test"
	// MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// DBの初期化
func dbInit() {
	db := gormConnect()
	// コネクション解放
	defer db.Close()
	db.AutoMigrate(&Blog{}) //構造体に基づいてテーブルを作成
	db.AutoMigrate(&User{})
}

// ユーザー登録処理
func createUser(mail string, password string) []error {
	passwordEncrypt, _ := crypto.PasswordEncrypt(password)
	db := gormConnect()
	defer db.Close()
	// Insert処理
	if err := db.Create(&User{Mail: mail, Password: passwordEncrypt}).GetErrors(); err != nil {
		return err
	}
	return nil

}

// ユーザーを一件取得
func getUser(mail string) User {
	db := gormConnect()
	var user User
	db.First(&user, "mail = ?", mail)
	db.Close()
	return user
}

// つぶやき登録処理
func createBlog(title string, content string, pictureUrl string) {
	db := gormConnect()
	defer db.Close()
	// Insert処理
	db.Create(&Blog{Title: title, Content: content, PictureUrl: pictureUrl})
}

// つぶやき更新
func updateBlog(id int, blogText string) {
	db := gormConnect()
	var blog Blog
	db.First(&blog, id)
	blog.Content = blogText
	db.Save(&blog)
	db.Close()
}

// つぶやき全件取得
func getAllBlogs() []Blog {
	db := gormConnect()

	defer db.Close()
	var blogs []Blog
	// FindでDB名を指定して取得した後、orderで登録順に並び替え
	db.Order("created_at desc").Find(&blogs)
	return blogs
}

// つぶやき一件取得
func getBlog(id int) Blog {
	db := gormConnect()
	var blog Blog
	db.First(&blog, id)
	db.Close()
	return blog
}

// つぶやき削除
func deleteBlog(id int) {
	db := gormConnect()
	var blog Blog
	db.First(&blog, id)
	db.Delete(&blog)
	db.Close()
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	dbInit()

	// 一覧
	router.GET("/", func(c *gin.Context) {
		blogs := getAllBlogs()
		c.HTML(200, "index.html", gin.H{"blogs": blogs})
	})

	// ユーザー登録画面
	router.GET("/signup", func(c *gin.Context) {

		c.HTML(200, "signup.html", gin.H{})
	})

	// ユーザー登録
	router.POST("/signup", func(c *gin.Context) {
		var form User
		// バリデーション処理
		if err := c.Bind(&form); err != nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
			c.Abort()
		} else {
			mail := c.PostForm("mail")
			password := c.PostForm("password")
			// 登録ユーザーが重複していた場合にはじく処理
			if err := createUser(mail, password); err != nil {
				c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
			}
			c.Redirect(302, "/")
		}
	})

	// ユーザーログイン画面
	router.GET("/login", func(c *gin.Context) {

		c.HTML(200, "login.html", gin.H{})
	})

	// ユーザーログイン
	router.POST("/login", func(c *gin.Context) {

		// DBから取得したユーザーパスワード(Hash)
		dbPassword := getUser(c.PostForm("mail")).Password
		log.Println(dbPassword)
		// フォームから取得したユーザーパスワード
		formPassword := c.PostForm("password")

		// ユーザーパスワードの比較
		if err := crypto.CompareHashAndPassword(dbPassword, formPassword); err != nil {
			log.Println("ログインできませんでした")
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
			c.Abort()
		} else {
			log.Println("ログインできました")
			c.Redirect(302, "/")
		}
	})

	// 記事新規作成
	router.POST("/new", func(c *gin.Context) {
		var form Blog
		// バリデーション処理
		if err := c.Bind(&form); err != nil {
			blogs := getAllBlogs()
			c.HTML(http.StatusBadRequest, "index.html", gin.H{"blogs": blogs, "err": err})
			c.Abort()
		} else {
			title := c.PostForm("title")
			content := c.PostForm("content")
			pictureUrl := c.PostForm("pictureUrl")
			createBlog(title, content, pictureUrl)
			c.Redirect(302, "/")
		}
	})

	// つぶやき詳細
	router.GET("/detail/:id", func(c *gin.Context) {
		n := c.Param("id")
		// パラメータから受け取った値をint化
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		blog := getBlog(id)
		c.HTML(200, "detail.html", gin.H{"blog": blog})
	})

	// 更新
	router.POST("/update/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		blog := c.PostForm("blog")
		updateBlog(id, blog)
		c.Redirect(302, "/")
	})

	// 削除確認
	router.GET("/delete_check/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		blog := getBlog(id)
		c.HTML(200, "delete.html", gin.H{"blog": blog})
	})

	// 削除
	router.POST("/delete/:id", func(c *gin.Context) {
		n := c.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		deleteBlog(id)
		c.Redirect(302, "/")

	})

	router.Run()
}
