package repoimpl

import (
	"bytes"
	"context"
	"database/sql"
	"html/template"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/userdm"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
	mydatabase "github.com/schwarzwald0906/My_Portfolio/src/core/infra/database"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/datamodel"
)

type UserRepoImpl struct {
	db *sqlx.DB
}

// データベースのコネクションを外側から渡せるようにすることでテストが容易になります。
//
// また、関数の戻り値をインタフェース型にして構造体をreturnすると型チェックが行われます。
// 構造体がインタフェースを満たしていない場合はコンパイルエラーになるのですぐに気付けて便利です。
func NewUserRepository(db *sql.DB) userdm.UserRepository {
	return &UserRepoImpl{
		db: &sqlx.DB{},
	}
}

// Create implements userdm.UserRepository
func (repo *UserRepoImpl) Create(ctx context.Context, user *userdm.User) error {
	//データベース接続
	repo.db = mydatabase.DbInit()

	// テンプレートをパースする
	tmpl, err := template.ParseFiles("create_user.sql")
	if err != nil {
		return err
	}

	// テンプレートを埋め込んだ結果を出力する
	var buf bytes.Buffer

	// テンプレートに値を代入する
	data := map[string]string{
		"UserId":    user.ID().String(),
		"Email":     user.Email().Value(),
		"Password":  user.Password().Value(),
		"CreatedAt": user.CreatedAt().Value().String(),
		"UpdatedAt": user.UpdatedAt().Value().String(),
	}

	tmpl.Execute(&buf, data)

	// 埋め込まれたSQLクエリを文字列として取得する
	sql := buf.String()

	err = repo.db.QueryRow(sql).Scan(&repo)
	if err != nil {
		log.Fatalln(err)
	}
	defer repo.db.Close()
	return nil

}

// FindByEmailID implements userdm.UserRepository
func (repo *UserRepoImpl) FindByEmailID(ctx context.Context, email vo.Email) (*userdm.User, error) {
	var scanuser datamodel.User
	var dmuser *userdm.User

	//データベース接続
	repo.db = mydatabase.DbInit()

	// テンプレートをパースする
	tmpl, err := template.ParseFiles("find_by_email.sql")
	if err != nil {
		return dmuser, err
	}

	// テンプレートに値を代入する
	data := map[string]string{
		"Email": email.Value(),
	}

	// テンプレートを埋め込んだ結果を出力する
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)

	// 埋め込まれたSQLクエリを文字列として取得する
	sql := buf.String()

	//単一行を返却するため、QueryRow,Scan
	err = repo.db.QueryRow(sql).Scan(&scanuser)
	if err != nil {
		return dmuser, err
	}
	// scanuserからdomainmuserへ型変換
	dmuser, err = userdm.Reconstruct(scanuser.ID(), scanuser.Email(), scanuser.Password(), scanuser.CreatedAt(), scanuser.UpdatedAt())

	if err != nil {
		return dmuser, err
	}
	defer repo.db.Close()
	return dmuser, nil

}

// FindByUserID implements userdm.UserRepository
func (repo *UserRepoImpl) FindByUserID(ctx context.Context, userId userdm.UserID) (*userdm.User, error) {
	var scanuser datamodel.User
	var dmuser *userdm.User

	//データベース接続
	repo.db = mydatabase.DbInit()

	// テンプレートをパースする
	tmpl, err := template.ParseFiles("find_by_user_id.sql")
	if err != nil {
		return dmuser, err
	}

	// テンプレートに値を代入する
	data := map[string]string{
		"UserId": userId.String(),
	}

	// テンプレートを埋め込んだ結果を出力する
	var buf bytes.Buffer
	tmpl.Execute(&buf, data)

	// 埋め込まれたSQLクエリを文字列として取得する
	sql := buf.String()

	//単一行を返却するため、QueryRow,Scan
	err = repo.db.QueryRow(sql).Scan(&scanuser)
	if err != nil {
		return dmuser, err
	}
	// scanuserからdomainmuserへ型変換
	dmuser, err = userdm.Reconstruct(scanuser.ID(), scanuser.Email(), scanuser.Password(), scanuser.CreatedAt(), scanuser.UpdatedAt())

	if err != nil {
		return dmuser, err
	}
	defer repo.db.Close()

	return dmuser, nil
}
