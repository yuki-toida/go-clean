package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/go-clean/src/adapter/repositories/repository_email"
	"github.com/yuki-toida/go-clean/src/adapter/repositories/repository_profile"
	"github.com/yuki-toida/go-clean/src/adapter/repositories/repository_user"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_email"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_profile"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_user"
	"github.com/yuki-toida/go-clean/src/registry/interfaces"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.Repository {
	return &repository{db: db}
}

func (r *repository) NewUserRepository() entity_user.Repository {
	return repository_user.New(r.db)
}

func (r *repository) NewProfileRepository() entity_profile.Repository {
	return repository_profile.New(r.db)
}

func (r *repository) NewEmailRepository() entity_email.Repository {
	return repository_email.New(r.db)
}
