package service

import "repository-hotel-booking/internal/app/model"

func (s *Service) GetAccounts(q *model.AccountQuery) ([]model.Account, error) {
	accounts, err := s.repo.AccountRepo.GetAccounts(q)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
