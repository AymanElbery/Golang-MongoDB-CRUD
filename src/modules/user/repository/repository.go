package repository

import (
	"gomongo/src/modules/user/model"
)

// UserRepository ...
type UserRepository interface {
	Save(*model.User) error
	Update(string, *model.User) error
	Delete(string) error
	FindByID(string) (*model.User, error)
	FindAll() (model.Users, error)
}
