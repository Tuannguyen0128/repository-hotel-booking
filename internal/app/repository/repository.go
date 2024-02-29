package repository

import (
	"database/sql"
	"repository-hotel-booking/internal/app/repository/account"
	"repository-hotel-booking/internal/app/repository/staff"
)

type (
	Repositories struct {
		AccountRepo *account.Repository
		StaffRepo   *staff.Repository
	}
)

func New(db *sql.DB) *Repositories {
	accountRepo := account.New(db)
	staffRepo := staff.New(db)

	return &Repositories{
		AccountRepo: accountRepo,
		StaffRepo:   staffRepo,
	}
}

//func (r *Repositories) Close() {
//	if r.pfDb != nil {
//		_ = r.pfDb.Close()
//	}
//	if r.txbDb != nil {
//		_ = r.txbDb.Close()
//	}
//}
