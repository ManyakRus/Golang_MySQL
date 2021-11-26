package usecase

import (
	"MySQL/domain"
)

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
}
