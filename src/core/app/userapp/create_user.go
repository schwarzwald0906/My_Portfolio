package userapp

import (
	"context"
	"time"

	"github.com/ymdd1/mytweet/src/core/domain/userdm"
	"github.com/ymdd1/mytweet/src/core/domain/vo"
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
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
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

	// createdAt, err := vo.NewCreatedAt(req.CreatedAt.Value())
	// if err != nil {
	// 	return nil, err
	// }

	// updatedAt, err := vo.NewUpdatedAt(req.UpdatedAt.Value())
	// if err != nil {
	// 	return nil, err
	// }

	// user, err := userdm.NewUser(userdm.NewUserID(), email, password, createdAt, updatedAt)
	// if err != nil {
	// 	return nil, err
	// }

	//入力値からドメインモデルを取得
	user, err := userdm.GenWhenCreate(email, password)
	if err != nil {
		return err
	}

	//上記で作成したuserをもとにINSERT処理を実行
	err = app.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
