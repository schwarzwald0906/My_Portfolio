package userdm

import (
	"context"

	"github.com/schwarzwald0906/My_Portfolio/src/core/domain/vo"
)

type IsDuplicatedByEmailDomainService struct {
	userRepository UserRepository
}

func NewIsDuplicatedByEmailDomainService(userRepo UserRepository) *IsDuplicatedByEmailDomainService {
	return &IsDuplicatedByEmailDomainService{userRepository: userRepo}
}

func (ds *IsDuplicatedByEmailDomainService) Exec(c context.Context, email vo.Email) (bool, error) {
	user, err := ds.userRepository.FindByEmailID(c, email)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
