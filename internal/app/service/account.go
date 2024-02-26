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

func (s *Service) UpdateAccount(a *model.Account) (*model.Account, *model.ErrInfo) {
	account, err := s.repo.AccountRepo.UpdateAccount(a)
	return account, err
}
func (s *Service) DeleteAccount(id string) (string, *model.ErrInfo) {
	result, err := s.repo.AccountRepo.DeleteAccount(id)
	return result, err
}
