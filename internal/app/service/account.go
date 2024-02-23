package service

import "repository-hotel-booking/internal/app/model"

func (s *Service) GetAccounts(q *model.AccountQuery) ([]model.Account, *model.ErrInfo) {
	accounts, err := s.repo.AccountRepo.GetAccounts(q)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (s *Service) AddAccount(a *model.Account) (string, *model.ErrInfo) {
	id, err := s.repo.AccountRepo.InsertAccount(a)
	if err != nil {
		return "", err
	}
	return id, nil
}
