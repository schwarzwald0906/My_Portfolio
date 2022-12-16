package repoimpl

import (
	"context"
	"embed"
	"log"

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

// Create implements userdm.UserRepository
func (repo *UserRepoImpl) Create(ctx context.Context, user *userdm.User) error {
	//go:embed embed/user/create_user.sql
	var usersql embed.FS
	tmpl, err := usersql.ReadFile("create_user.sql")
	if err != nil {
		return err
	}
	p := string(tmpl)

	// パラメータを渡してクエリを実行
	if _, err = repo.db.Exec(p, user.ID(), user.Email(), user.Password(), user.CreatedAt(), user.UpdatedAt()); err != nil {
		log.Fatalln(err)
	}
	return nil

}

// FindByEmailID implements userdm.UserRepository
func (repo *UserRepoImpl) FindByEmailID(ctx context.Context, email vo.Email) (*userdm.User, error) {
	//go:embed embed/user/find_by_email.sql
	var usersql embed.FS
	tmpl, err := usersql.ReadFile("find_by_email.sql")
	if err != nil {
		return nil, err
	}

	// パラメータを渡してクエリを実行
	var scanuser datamodel.User
	p := string(tmpl)
	err = repo.db.QueryRow(p, email).Scan(&scanuser)
	if err != nil {
		return nil, err
	}

	// scanuserからdmuserへ型変換
	dmuser, err := userdm.Reconstruct(scanuser.ID, scanuser.Email, scanuser.Password, scanuser.CreatedAt, scanuser.UpdatedAt)
	if err != nil {
		return dmuser, err
	}
	defer repo.db.Close()
	return dmuser, nil
}

// FindByUserID implements userdm.UserRepository
func (repo *UserRepoImpl) FindByUserID(ctx context.Context, userId userdm.UserID) (*userdm.User, error) {
	//go:embed embed/user/find_by_user_id.sql
	var usersql embed.FS
	tmpl, err := usersql.ReadFile("find_by_user_id.sql")
	if err != nil {
		return nil, err
	}

	// パラメータを渡してクエリを実行
	var scanuser datamodel.User
	p := string(tmpl)
	err = repo.db.QueryRow(p, userId).Scan(&scanuser)
	if err != nil {
		return nil, err
	}

	// scanuserからdmuserへ型変換
	dmuser, err := userdm.Reconstruct(scanuser.ID, scanuser.Email, scanuser.Password, scanuser.CreatedAt, scanuser.UpdatedAt)
	if err != nil {
		return dmuser, err
	}
	return dmuser, nil
}
