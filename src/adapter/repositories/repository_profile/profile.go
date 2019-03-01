package repository_profile

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/go-clean/src/domain/entities"
	"github.com/yuki-toida/go-clean/src/domain/entities/entity_profile"
)

type Profile struct {
	entities.Model
	Name string `gorm:"not null"`
}

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db: db}
}

func ToEntity(profile *Profile) *entity_profile.Profile {
	return &entity_profile.Profile{
		Model: profile.Model,
		Name:  profile.Name,
	}
}

func (r *repository) Create(name string) (*entity_profile.Profile, error) {
	dbProfile := Profile{Name: name}
	if err := r.db.Create(&dbProfile).Error; err != nil {
		return nil, err
	}
	return ToEntity(&dbProfile), nil
}
