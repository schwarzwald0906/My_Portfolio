package userdm

import (
	"context"
)

type IsDuplicatedByUserIdDomainService struct {
	userRepository UserRepository
}

func NewUserDomainService(userRepo UserRepository) *IsDuplicatedByUserIdDomainService {
	return &IsDuplicatedByUserIdDomainService{userRepository: userRepo}
}

func (ds *IsDuplicatedByUserIdDomainService) IsExists(ctx context.Context, userID UserID) (bool, error) {
	user, err := ds.userRepository.FindByUserID(ctx, userID)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
