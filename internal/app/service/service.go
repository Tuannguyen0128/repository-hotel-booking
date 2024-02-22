package service

import (
	"repository-hotel-booking/internal/app/repository"
)

type Service struct {
	repo *repository.Repositories
}

func NewService(repo *repository.Repositories) *Service {
	return &Service{repo: repo}
}
