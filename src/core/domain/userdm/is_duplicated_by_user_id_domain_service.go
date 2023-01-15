package userdm

import "context"

type IsDuplicatedByUserIDDomainService struct {
	userRepository UserRepository
}

func NewIsDuplicatedByUserIDDomainService(userRepo UserRepository) *IsDuplicatedByUserIDDomainService {
	return &IsDuplicatedByUserIDDomainService{userRepository: userRepo}
}

func (ds *IsDuplicatedByUserIDDomainService) Exec(c context.Context, userID UserID) (bool, error) {
	user, err := ds.userRepository.FindByUserID(c, userID)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
