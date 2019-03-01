package interactor_email

import (
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_email"
)

type interactor struct {
	emailRepository entity_email.Repository
}

func New(er entity_email.Repository) *interactor {
	return &interactor{emailRepository: er}
}

func (i *interactor) Create(userID uint64, emailAddr string) (*entity_email.Email, error) {
	ee := entity_email.New(i.emailRepository)
	email, err := ee.Create(userID, emailAddr)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (i *interactor) Update(emailID uint64, emailAddr string) (*entity_email.Email, error) {
	ee := entity_email.New(i.emailRepository)
	email, err := ee.Update(emailID, emailAddr)
	if err != nil {
		return nil, err
	}
	return email, nil
}
