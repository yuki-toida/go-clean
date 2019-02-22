package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/go-clean/src/adapter/repositories"
	application "github.com/yuki-toida/go-clean/src/application/repositories"
	"github.com/yuki-toida/go-clean/src/registry/interfaces"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.Repository {
	return &repository{db: db}
}

func (r *repository) NewUserRepository() application.UserRepository {
	return repositories.NewUserRepository(r.db)
}

func (r *repository) NewProfileRepository() application.ProfileRepository {
	return repositories.NewProfileRepository(r.db)
}

func (r *repository) NewEmailRepository() application.EmailRepository {
	return repositories.NewEmailRepository(r.db)
}
