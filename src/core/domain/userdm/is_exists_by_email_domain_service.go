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

func (ds *IsDuplicatedByEmailDomainService) IsExists(ctx context.Context, email vo.Email) bool {
	user, err := ds.userRepository.FindByEmailID(ctx, email)
	return !(err != nil || user == nil)
}
