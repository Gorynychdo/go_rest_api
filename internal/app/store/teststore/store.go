package teststore

import (
	"github.com/Gorynychdo/go_rest_api/internal/app/model"
	"github.com/Gorynychdo/go_rest_api/internal/app/store"
	_ "github.com/lib/pq"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
