package interactors

import (
	"github.com/yuki-toida/go-clean/src/application/repositories"
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type emailInteractor struct {
	emailRepository repositories.EmailRepository
}

func NewEmailInteractor(er repositories.EmailRepository) *emailInteractor {
	return &emailInteractor{emailRepository: er}
}

func (i *emailInteractor) Create(userID uint64, email string) (*entities.Email, error) {
	e, err := i.emailRepository.Create(userID, email)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (i *emailInteractor) Update(emailID uint64, email string) (*entities.Email, error) {
	e, err := i.emailRepository.Update(emailID, email)
	if err != nil {
		return nil, err
	}
	return e, nil
}
