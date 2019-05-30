package repository

import (
	"github.com/jinzhu/gorm"
)

type user interface{}

//Repository interface
type Repository interface {
	CreateUser(req user) (*user, error)
}

//UserRepository struct
type UserRepository struct {
	Db *gorm.DB
}

//CreateUser to table
func (repo *UserRepository) CreateUser() (*user, error) {
	return nil, nil
}
