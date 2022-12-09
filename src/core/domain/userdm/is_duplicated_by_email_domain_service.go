package userdm

import (
	"context"

	"github.com/ymdd1/mytweet/src/core/domain/vo"
)

type IsDuplicatedByEmailDomainService struct {
	userRepository UserRepository
}

func NewEmailDomainService(userRepo UserRepository) *IsDuplicatedByEmailDomainService {
	return &IsDuplicatedByEmailDomainService{userRepository: userRepo}
}

func (ds *IsDuplicatedByEmailDomainService) Exec(ctx context.Context, email vo.Email) (bool, error) {
	user, err := ds.userRepository.FindByEmailID(ctx, email)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
