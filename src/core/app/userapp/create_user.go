package userapp

import (
	"context"

	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/userdm"
	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type CreateUserApp struct {
	userRepository userdm.UserRepository
}

func NewCreateUserApp(userRepo userdm.UserRepository) *CreateUserApp {
	return &CreateUserApp{
		userRepository: userRepo,
	}
}

type CreateUserRequest struct {
	Email    string
	Password string
}

func (app *CreateUserApp) Exec(c context.Context, req *CreateUserRequest) error {

	email, err := vo.NewEmail(c, req.Email)
	if err != nil {
		return err
	}
	password, err := vo.NewPassword(c, req.Password)
	if err != nil {
		return err
	}
	//入力値からドメインモデルを取得
	user, err := userdm.GenWhenCreate(email, password)
	if err != nil {
		return err
	}
	//上記で作成したuserをもとにINSERT処理を実行
	return app.userRepository.Create(c, user)

}
