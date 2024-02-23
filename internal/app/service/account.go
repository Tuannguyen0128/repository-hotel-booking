package service

import "repository-hotel-booking/internal/app/model"

func (s *Service) GetAccounts(q *model.AccountQuery) ([]model.Account, *model.ErrInfo) {
	accounts, err := s.repo.AccountRepo.GetAccounts(q)
	return accounts, err
}

func (s *Service) AddAccount(a *model.Account) (string, *model.ErrInfo) {
	id, err := s.repo.AccountRepo.InsertAccount(a)
	return id, err
}
