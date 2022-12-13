package repoimpl

import (
	"context"
	"database/sql"
	"embed"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/userdm"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
	mydatabase "github.com/schwarzwald0906/My_Portfolio/src/core/infra/database"
	"github.com/schwarzwald0906/My_Portfolio/src/core/infra/datamodel"
)

//go:embed embed/*.sql
var emb embed.FS

type UserRepoImpl struct {
	db *sqlx.DB
}

// データベースのコネクションを外側から渡せるようにすることでテストが容易になります。
//
// また、関数の戻り値をインタフェース型にして構造体をreturnすると型チェックが行われます。
// 構造体がインタフェースを満たしていない場合はコンパイルエラーになるのですぐに気付けて便利です。
func NewUserRepository(db *sql.DB) userdm.UserRepository {
	return &UserRepoImpl{
		db: &sqlx.DB{
			DB:     db,
			Mapper: &reflectx.Mapper{},
		},
	}
}

// Create implements userdm.UserRepository
func (repo *UserRepoImpl) Create(ctx context.Context, user *userdm.User) error {
	//データベース接続
	repo.db = mydatabase.DbInit()

	// go-embed
	tmpl, err := emb.ReadFile("sql/create_user.sql")
	if err != nil {
		return err
	}
	p := string(tmpl)

	// プリペアドステートメントを作成
	stmt, err := repo.db.Prepare(p)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	id := user.ID()
	email := user.Email()
	password := user.Password()
	createdat := user.CreatedAt()
	updatedat := user.UpdatedAt()

	// パラメータを渡してクエリを実行
	err = stmt.QueryRow(id, email, password, createdat, updatedat).Scan(&repo)
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

	// go-embed
	tmpl, err := emb.ReadFile("sql/find_by_email.sql")
	if err != nil {
		return dmuser, err
	}
	p := string(tmpl)

	// プリペアドステートメントを作成
	stmt, err := repo.db.Prepare(p)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// パラメータを渡してクエリを実行
	err = stmt.QueryRow(email).Scan(&scanuser)
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

	// データベース接続
	repo.db = mydatabase.DbInit()

	// go-embed
	tmpl, err := emb.ReadFile("sql/find_by_user_id.sql")
	if err != nil {
		return dmuser, err
	}
	p := string(tmpl)

	// プリペアドステートメントを作成
	stmt, err := repo.db.Prepare(p)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// パラメータを渡してクエリを実行
	err = stmt.QueryRow(userId).Scan(&scanuser)
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
