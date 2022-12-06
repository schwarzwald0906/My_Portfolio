package userdm

import "context"

type EmailDomainService struct {
	emailRepository EmailRepository
}

func NewEmailDomainService(emailRepo EmailRepository) *EmailDomainService {
	return &EmailDomainService{emailRepository: emailRepo}
}

func (ds *EmailDomainService) IsExists(ctx context.Context, email string) (bool, error) {
	err := ds.emailRepository.FindByEmailID(ctx, email)
	if err != nil {
		return false, err
	}
	return true, nil
}
