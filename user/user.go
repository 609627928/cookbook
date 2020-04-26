package user

import "cookbook/model"

type Store interface {
	GetByID(uint) (*model.User, error)
	GetByEmail(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
}
