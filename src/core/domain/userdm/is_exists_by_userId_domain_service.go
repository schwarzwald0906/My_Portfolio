package userdm

import (
	"context"
)

type UserDomainService struct {
	userRepository UserRepository
}

func NewUserDomainService(userRepo UserRepository) *UserDomainService {
	return &UserDomainService{userRepository: userRepo}
}

func (ds *UserDomainService) IsExists(ctx context.Context, userID UserID) (bool, error) {
	user, err := ds.userRepository.FindByUserID(ctx, userID)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
