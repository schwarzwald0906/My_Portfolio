package userdm

import "context"

type UserDomainService struct {
	userRepository UserRepository
}

func NewUserDomainService(userRepo UserRepository) *UserDomainService {
	return &UserDomainService{userRepository: userRepo}
}

func (ds *UserDomainService) IsExists(ctx context.Context, userID string) (bool, error) {
	_, err := ds.userRepository.FindByID(ctx, userID)
	if err != nil {
		return false, err
	}
	return true, nil
}
