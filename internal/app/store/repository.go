package store

import "github.com/Gorynychdo/go_rest_api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
