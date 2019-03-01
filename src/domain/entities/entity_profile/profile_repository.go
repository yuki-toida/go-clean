package entity_profile

type Repository interface {
	Create(name string) (*Profile, error)
}
