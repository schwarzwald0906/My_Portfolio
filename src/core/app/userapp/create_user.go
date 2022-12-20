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
	// CreatedAt time.Time
	// UpdatedAt time.Time
}

// type CreateUserResponse struct {
// 	ID vo.UserId
// }

func (app *CreateUserApp) Exec(ctx context.Context, req *CreateUserRequest) error {
	email, err := vo.NewEmail(req.Email)
	if err != nil {
		return err
	}
	password, err := vo.NewPassword(req.Password)
	if err != nil {
		return err
	}

	//入力値からドメインモデルを取得
	user, err := userdm.GenWhenCreate(email, password)
	if err != nil {
		return err
	}

	//上記で作成したuserをもとにINSERT処理を実行
	return app.userRepository.Create(ctx, user)
}
