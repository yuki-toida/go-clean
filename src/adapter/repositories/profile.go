package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type Profile struct {
	entities.Model
	Name string `gorm:"not null"`
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *profileRepository {
	return &profileRepository{db: db}
}

func toProfileEntity(profile *Profile) *entities.Profile {
	return &entities.Profile{
		Model: profile.Model,
		Name:  profile.Name,
	}
}

func (r *profileRepository) Create(name string) (*entities.Profile, error) {
	dbProfile := Profile{Name: name}
	if err := r.db.Create(&dbProfile).Error; err != nil {
		return nil, err
	}
	return toProfileEntity(&dbProfile), nil
}
