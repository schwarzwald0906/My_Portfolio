package userdm

import (
	"context"
)

type IsDuplicatedByUserIDDomainService struct {
	userRepository UserRepository
}

func NewIsDuplicatedByUserIDDomainService(userRepo UserRepository) *IsDuplicatedByUserIDDomainService {
	return &IsDuplicatedByUserIDDomainService{userRepository: userRepo}
}

func (ds *IsDuplicatedByUserIDDomainService) Exec(ctx context.Context, userID UserID) (bool, error) {
	user, err := ds.userRepository.FindByUserID(ctx, userID)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
