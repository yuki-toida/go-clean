package usecase_email

import (
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_email"
)

type UseCase interface {
	Create(userID uint64, emailAddr string) (*entity_email.Email, error)
	Update(emailID uint64, emailAddr string) (*entity_email.Email, error)
}
