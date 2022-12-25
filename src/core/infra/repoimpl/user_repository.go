package repoimpl

import (
	"context"
	"embed"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/userdm"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/datamodel"
)

type UserRepoImpl struct {
	db *sqlx.DB
}

// データベースのコネクションを外側から渡せるようにすることでテストが容易になる。
// また、関数の戻り値をインタフェース型にして構造体をreturnすると型チェックが行われる。
// 構造体がインタフェースを満たしていない場合はコンパイルエラーになるのですぐに気付けて便利。
func NewUserRepository(db *sqlx.DB) userdm.UserRepository {
	return &UserRepoImpl{
		db: db,
	}
}

//go:embed embed/user/create_user.sql
var createUserSQL embed.FS

// Create implements userdm.UserRepository
func (repo *UserRepoImpl) Create(ctx context.Context, user *userdm.User) error {
	fmt.Println("UserRepositoryのcreateまで到達")
	tmpl, err := createUserSQL.ReadFile("embed/user/create_user.sql")
	if err != nil {
		fmt.Printf("ReadERR")
		return err
	}
	// パラメータを渡してクエリを実行
	fmt.Println(user.ID().String())
	fmt.Println(user.Email().Value())
	fmt.Println(user.Password().Value())
	fmt.Println(user.CreatedAt().Value().Format("2006-01-02 15:04:05"))
	fmt.Println(user.UpdatedAt().Value().Format("2006-01-02 15:04:05"))
	if _, err = repo.db.Exec(
		string(tmpl),
		user.ID().String(),
		user.Email().Value(),
		user.Password().Value(),
		user.CreatedAt().Value().Format("2006-01-02 15:04:05"),
		user.UpdatedAt().Value().Format("2006-01-02 15:04:05")); err != nil {
		// log.Fatalln(err)
		fmt.Printf("ExecERR")
		return err
	}
	// if _, err = repo.db.Exec(string(tmpl), user.ID(), user.Email(), user.Password(), user.CreatedAt(), user.UpdatedAt()); err != nil {
	// 	// log.Fatalln(err)
	// 	fmt.Printf("ExecERR")
	// 	return err
	// }
	return nil

}

//go:embed embed/user/find_by_email.sql
var findByEmailSQL embed.FS

// FindByEmailID implements userdm.UserRepository
func (repo *UserRepoImpl) FindByEmailID(ctx context.Context, email vo.Email) (*userdm.User, error) {
	tmpl, err := findByEmailSQL.ReadFile("embed/user/find_by_email.sql")
	if err != nil {
		return nil, err
	}

	// パラメータを渡してクエリを実行
	var scanUser datamodel.User
	err = repo.db.QueryRow(string(tmpl), email).Scan(&scanUser)
	if err != nil {
		return nil, err
	}

	// scanUserからdmuserへ型変換
	return userdm.Reconstruct(scanUser.ID, scanUser.Email, scanUser.Password, scanUser.CreatedAt, scanUser.UpdatedAt)
}

//go:embed embed/user/find_by_user_id.sql
var findByUserSQL embed.FS

// FindByUserID implements userdm.UserRepository
func (repo *UserRepoImpl) FindByUserID(ctx context.Context, userId userdm.UserID) (*userdm.User, error) {
	tmpl, err := findByUserSQL.ReadFile("embed/user/find_by_user_id.sql")
	if err != nil {
		return nil, err
	}

	// パラメータを渡してクエリを実行
	var scanUser datamodel.User
	err = repo.db.QueryRow(string(tmpl), userId).Scan(&scanUser)
	if err != nil {
		return nil, err
	}

	// scanUserからdmuserへ型変換
	return userdm.Reconstruct(scanUser.ID, scanUser.Email, scanUser.Password, scanUser.CreatedAt, scanUser.UpdatedAt)
}
