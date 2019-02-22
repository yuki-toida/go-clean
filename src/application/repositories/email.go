package repositories

import (
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type EmailRepository interface {
	Create(userID uint64, email string) (*entities.Email, error)
	Update(emailID uint64, email string) (*entities.Email, error)
}
