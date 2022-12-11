package repoimpl

import (
	"context"
	"database/sql"
	"log"

	"github.com/ymdd1/mytweet/src/core/domain/userdm"
	"github.com/ymdd1/mytweet/src/core/domain/vo"
	myDatabase "github.com/ymdd1/mytweet/src/core/infra/database"
)

type UserRepoImpl struct {
	ID        vo.ID        `db:"id"`
	Email     vo.Email     `db:"email"`
	Password  vo.Password  `db:"password"`
	CreatedAt vo.CreatedAt `db:"created_at"`
	UpdatedAt vo.UpdatedAt `db:"updated_at"`
}

// データベースのコネクションを外側から渡せるようにすることでテストが容易になります。
//
// また、関数の戻り値をインタフェース型にして構造体をreturnすると型チェックが行われます。
// 構造体がインタフェースを満たしていない場合はコンパイルエラーになるのですぐに気付けて便利です。
func NewUserRepository(db *sql.DB) userdm.UserRepository {
	return &UserRepoImpl{
		ID:        "",
		Email:     "",
		Password:  "",
		CreatedAt: vo.CreatedAt{},
		UpdatedAt: vo.UpdatedAt{},
	}
}

// Create implements userdm.UserRepository
func (repo *UserRepoImpl) Create(ctx context.Context, user *userdm.User) error {
	//データベース接続
	db := myDatabase.DbInit()

	//SQL文定義
	sql := `INSERT INTO user
				(id, email, password,created_at, updated_at)
			VALUES
				(:id,:email,:password,:created_at,:updated_at);`

	err := db.QueryRow(sql, repo.ID, repo.Email, repo.Password, repo.CreatedAt, repo.UpdatedAt).Scan(&repo)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	return nil

}

// FindByEmailID implements userdm.UserRepository
func (repo *UserRepoImpl) FindByEmailID(ctx context.Context, email vo.Email) (*userdm.User, error) {
	//データベース接続
	db := myDatabase.DbInit()

	//SQL文定義
	sql := `SELECT 
				*
			FROM
				user  
			WHERE
				user.email = ?;`

	var user *userdm.User
	rows, err := db.Queryx(sql, repo.Email)
	if err != nil {
		return user, err
	}
	if err := rows.StructScan(user); err != nil {
		return user, err
	}
	rows.StructScan(user)
	defer db.Close()
	return user, nil

}

// FindByUserID implements userdm.UserRepository
func (repo *UserRepoImpl) FindByUserID(ctx context.Context, userId userdm.UserID) (*userdm.User, error) {
	//データベース接続
	db := myDatabase.DbInit()

	//SQL文定義
	sql := `SELECT 
				*
			FROM
				user  
			WHERE
				user.userId = ?;`

	var user *userdm.User
	rows, err := db.Queryx(sql, repo.ID)
	if err != nil {
		return user, err
	}
	if err := rows.StructScan(user); err != nil {
		return user, err
	}
	rows.StructScan(user)
	defer db.Close()
	return user, nil
}
