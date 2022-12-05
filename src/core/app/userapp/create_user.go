package userapp

import (
	"context"

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
	Email      vo.Email
	Password   vo.Password
	Created_at vo.Created_at
	Updated_at vo.Updated_at
}

type CreateUserResponse struct {
	ID string
}

func (app *CreateUserApp) Exec(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	email, err := vo.NewEmail(req.Email.Value())
	if err != nil {
		return nil, err
	}
	password, err := vo.NewPassword(req.Password.Value())
	if err != nil {
		return nil, err
	}

	created_at, err := vo.NewCreated_at(req.Created_at.Value())
	if err != nil {
		return nil, err
	}

	updated_at, err := vo.NewUpdated_at(req.Updated_at.Value())
	if err != nil {
		return nil, err
	}

	user, err := userdm.NewUser(userdm.NewUserID(), email, password, created_at, updated_at)
	if err != nil {
		return nil, err
	}

	createdUser, err := app.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{ID: createdUser.ID().String()}, nil
}
