package service

import "repository-hotel-booking/internal/app/model"

func (s *Service) GetStaffs(q *model.StaffQuery) ([]model.Staff, *model.ErrInfo) {
	staffs, err := s.repo.StaffRepo.GetStaffs(q)
	return staffs, err
}

func (s *Service) AddStaff(st *model.Staff) (string, *model.ErrInfo) {
	id, err := s.repo.StaffRepo.InsertStaff(st)
	return id, err
}

func (s *Service) UpdateStaff(st *model.Staff) (*model.Staff, *model.ErrInfo) {
	staff, err := s.repo.StaffRepo.UpdateStaff(st)
	return staff, err
}

func (s *Service) DeleteStaff(id string) (string, *model.ErrInfo) {
	result, err := s.repo.StaffRepo.DeleteStaff(id)
	return result, err
}
