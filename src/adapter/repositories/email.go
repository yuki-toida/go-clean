package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/go-clean/src/domain/entities"
)

type Email struct {
	entities.Model
	Email  string `gorm:"not null"`
	UserID uint64 `gorm:"not null;index"`
}

type emailRepository struct {
	db *gorm.DB
}

func NewEmailRepository(db *gorm.DB) *emailRepository {
	return &emailRepository{db: db}
}

func toEmailEntity(email Email) *entities.Email {
	return &entities.Email{
		Model:  email.Model,
		Email:  email.Email,
		UserID: email.UserID,
	}
}

func (r *emailRepository) first(emailID uint64) (*Email, error) {
	email := Email{Model: entities.Model{ID: emailID}}
	if err := r.db.First(&email).Error; err != nil {
		return nil, err
	}
	return &email, nil
}

func (r *emailRepository) Create(userID uint64, email string) (*entities.Email, error) {
	e := Email{Email: email, UserID: userID}
	if err := r.db.Create(&e).Error; err != nil {
		return nil, err
	}
	return toEmailEntity(e), nil
}

func (r *emailRepository) Update(emailID uint64, email string) (*entities.Email, error) {
	e, err := r.first(emailID)
	if err != nil {
		return nil, err
	}
	e.Email = email
	if err := r.db.Save(e).Error; err != nil {
		return nil, err
	}
	return toEmailEntity(*e), nil
}
