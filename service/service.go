package service

import (
	"apipost/repositories/orders_repo"
	"apipost/repositories/users_repo"
)

type Service struct {
	Repositories Repositories
}

type Repositories struct {
	Users  *users_repo.Repo
	Orders *orders_repo.Repo
}

func New(repos Repositories) *Service {
	return &Service{
		Repositories: repos,
	}
}
