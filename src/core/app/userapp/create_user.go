package userapp

import (
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
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID string
}

func (app *CreateUserApp) Exec(req *CreateUserRequest) (*CreateUserResponse, error) {
	// email, err := vo.NewEmail(req.Email)
	// if err != nil {
	// 	return nil, err
	// }
	password, err := vo.NewPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user, err := userdm.NewUser(userdm.NewUserID(), req.Email, password)
	if err != nil {
		return nil, err
	}

	createdUser, err := app.userRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{ID: createdUser.ID().String()}, nil
}
