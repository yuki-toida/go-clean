package interfaces

import (
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_email"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_profile"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_user"
)

type Repository interface {
	NewUserRepository() entity_user.Repository
	NewProfileRepository() entity_profile.Repository
	NewEmailRepository() entity_email.Repository
}
