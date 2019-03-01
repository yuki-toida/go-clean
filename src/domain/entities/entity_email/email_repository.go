package entity_email

type Repository interface {
	Create(userID uint64, emailAddr string) (*Email, error)
	Update(emailID uint64, emailAddr string) (*Email, error)
}
